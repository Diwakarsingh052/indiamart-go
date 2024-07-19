package main

import (
	"log/slog"
	"os"
)

func setupLog() {
	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		//AddSource: true: This will cause the source file and line number of the log message to be included in the output
		AddSource: true,
	})

	logger := slog.New(logHandler)
	//SetDefault makes l the default Logger. in our case we would be doing structured logging
	slog.SetDefault(logger)
}
