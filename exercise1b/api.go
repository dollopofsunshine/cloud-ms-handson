package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func AllNotes(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "All Notes!")
}


func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Home).Methods("GET")
	r.HandleFunc("/notes", AllNotes).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}