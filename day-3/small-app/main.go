package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"small-app/handlers"
	"time"
)

func main() {
	setupLog()
	err := startApp()
	if err != nil {
		panic(err)
	}

}

func startApp() error {
	api := http.Server{
		Addr:         ":8000",
		ReadTimeout:  8000 * time.Second,
		WriteTimeout: 800 * time.Second,
		IdleTimeout:  800 * time.Second,
		Handler:      handlers.API(),
	}

	//shutdown channel intercepts ctrl+c signals
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)

	// channel to store any errors while setting up the service
	serverErrors := make(chan error, 1)
	go func() {
		serverErrors <- api.ListenAndServe()
	}()

	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error %w", err)
	case <-shutdown:
		fmt.Println("graceful shutdown")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		//Shutdown gracefully shuts down the server without interrupting any active connections.
		//Shutdown works by first closing all open listeners, then closing all idle connections,

		err := api.Shutdown(ctx)
		if err != nil {
			err := api.Close()
			if err != nil {
				return fmt.Errorf("could not stop server gracefully %w", err)
			}
		}
	}

	return nil

}
