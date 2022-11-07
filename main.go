package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func initializeRouter() {

	r := mux.NewRouter()

	r.HandleFunc("/login", Login).Methods("POST")

	r.HandleFunc("/verification", Verificatio).Methods("POST")

	log.Fatal(http.ListenAndServe(":9000", r))

}
