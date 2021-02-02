package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//urlShortener decodes url and matches end key with map
func urlShortener(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	urlName := params["id"]
	log.Println(urlName)
	w.Header().Set("Content-Type", "application/json")

	if value, ok := urlMap[urlName]; ok {
		log.Println("Redirecting to", urlName)
		http.Redirect(w, r, value, 302)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, `{"error": "404: URL Not Found"}`)

}
