package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("hello world!")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello", helloWorld).Methods("GET")
	log.Fatal(http.ListenAndServe(":6000", router))
}