package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"test/entity"

	"github.com/gorilla/mux"
)

var userData = entity.Users{
	{
		ID:          "0",
		Name:        "Aldo",
		Description: "Hello World!",
	},
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func getAll(responseWriter http.ResponseWriter, request *http.Request) {
	json.NewEncoder(responseWriter).Encode(userData)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/users", getAll).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", router))
}
