package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)



func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", CreateUserEndpoint).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}
