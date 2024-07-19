package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// download a module using go get
//go get github.com/gorilla/mux

func main() {
	r := mux.NewRouter().StrictSlash(true)

	// Create subrouter for posts
	postsRouter := r.PathPrefix("/posts").Subrouter()
	postsRouter.HandleFunc("/", AllPosts)
	postsRouter.HandleFunc("/{id}", GetPost)

	// Create subrouter for users
	usersRouter := r.PathPrefix("/user").Subrouter()
	usersRouter.HandleFunc("/", AllUsers).Methods(http.MethodGet)
	usersRouter.HandleFunc("/{id}", GetUser).Methods(http.MethodPost)
	http.ListenAndServe(":8000", r)

}

// Handler functions
func AllPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All Posts")
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	//fetching id variable from the url
	id := mux.Vars(r)["id"]
	fmt.Fprintf(w, "Post with ID: %s", id)
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All Users")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	//fetching id variable from the url
	id := mux.Vars(r)["id"]
	fmt.Fprintf(w, "User with ID: %s", id)
}
