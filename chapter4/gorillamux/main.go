package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get")
}

func PostNoteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("post")
}

func PutNoteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("put")
}

func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete")
}

func main() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/notes", GetNoteHandler).Methods("GET")
	r.HandleFunc("/api/notes", PostNoteHandler).Methods("POST")
	r.HandleFunc("/api/notes/{id}", PutNoteHandler).Methods("PUT")
	r.HandleFunc("/api/notes/{id}", DeleteNoteHandler).Methods("DELETE")

	server := &http.Server{
		Addr:    ":8080",
		Handler: r}
	server.ListenAndServe()
}
