package main

// import "./dao"

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// AllMoviesEndPoint : fetch all stored movie items
func AllMoviesEndPoint(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "not implemented yet")
}

// FindMovieEndPoint : Find a certain movie item
func FindMovieEndPoint(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "not implemented yet")
}

// CreateMovieEndPoint : creating a new movie item
func CreateMovieEndPoint(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "not implemented yet")
}

// UpdateMovieEndPoint : editing movie details
func UpdateMovieEndPoint(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "not implemented yet")
}

// DeleteMovieEndPoint : delete a stored movie item and its details
func DeleteMovieEndPoint(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "not implemented yet")
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/movies", AllMoviesEndPoint).Methods("GET")
	r.HandleFunc("/movies", CreateMovieEndPoint).Methods("POST")
	r.HandleFunc("/movies", UpdateMovieEndPoint).Methods("PUT")
	r.HandleFunc("/movies", DeleteMovieEndPoint).Methods("DELETE")
	r.HandleFunc("/movies/{id}", FindMovieEndPoint).Methods("GET")

	if err := http.ListenAndServe(":1234", r); err != nil {
		log.Fatal(err)
	}
}
