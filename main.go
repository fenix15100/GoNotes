package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"GoNotes/Api"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/people", Api.GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", Api.GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", Api.CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", Api.DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))

}


