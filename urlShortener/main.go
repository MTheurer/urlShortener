package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type urlFormat struct {
	Url string `json:"URL"`
}

var urlMap = make(map[string]string)

// HomePage func to return Homepage text
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my ShortyResty :)")
}

func main() {

	urlMap["ASDFASDF"] = "http://example.com/verylonguselessURLthatdoesnotseemtoend/123"

	HandleRequests()
}

func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", HomePage).Methods("GET")
	myRouter.HandleFunc("/shorten/{id}", urlShortener).Methods("GET")
	myRouter.HandleFunc("/create", createKey).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
