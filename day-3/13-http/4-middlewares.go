package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/home", Mid(LogMid(HomePage)))
	http.ListenAndServe(":8080", nil)
}

func Mid(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware invoked")
		next(w, r)
	}
}

func LogMid(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("middleware is called")
		log.Println("doing middleware specific things")
		log.Println(r.Method)
		next(w, r)
		log.Println(r.URL)
	}

}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("home page invoked")
	go panic("anything in panic")
	fmt.Fprintln(w, "this is my home")
}
