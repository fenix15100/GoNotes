package main

import (
	"GoNotes/Api"
	"GoNotes/Services"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	Services.GetDatabase(true)

	router := mux.NewRouter()
	router.HandleFunc("/people", Api.GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", Api.GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", Api.CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", Api.DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))

}


