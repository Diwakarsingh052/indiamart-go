package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	// route with path and query parameter
	r.HandleFunc("/item/{id}", FetchById)
	// regex to handle post names like : go-101 , learn-go
	r.HandleFunc("/post/{postSlug:[a-z0-9-]+}", getArticleBySlug)
	// route with query parameter
	r.HandleFunc("/search", SearchUser) // ?=123
	http.ListenAndServe(":8080", r)
}

func getArticleBySlug(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	articleSlug := vars["postSlug"]
	w.Write([]byte("Post Slug: " + articleSlug))
}

func FetchById(w http.ResponseWriter, r *http.Request) {
	// get path parameter
	vars := mux.Vars(r)
	id := vars["id"]

	// get query parameter
	queryParams := r.URL.Query()
	val := queryParams.Get("value")

	w.Write([]byte("ID: " + id + ", Query Value: " + val))
}

func SearchUser(w http.ResponseWriter, r *http.Request) {
	// get query parameter
	fmt.Println("hey")
	queryParams := r.URL.Query()
	val := queryParams.Get("name")

	w.Write([]byte("Query Value: " + val))
}

//curl http://localhost:8080/my-article
//curl http://localhost:8080/another-article
//curl http://localhost:8080/article-with
//Match any sequence of one or more characters that are either a lowercase letter or a hyphen
