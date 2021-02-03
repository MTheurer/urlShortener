package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
)

//RandomString generates a random string of A-Z chars
func RandomString(len int) string {
	bytes := make([]byte, len)

	for i := 0; i < len; i++ {
		//A=65 and Z = 65+25
		bytes[i] = byte(65 + rand.Intn(25))
	}
	return string(bytes)
}

//createKey decodes the recieved Json and detects errors
func createKey(w http.ResponseWriter, r *http.Request) {
	log.Println("starting creatKey server")
	w.Header().Set("Content-Type", "application/json")

	var newUrl urlFormat

	fmt.Println("decoding newKey")
	err := json.NewDecoder(r.Body).Decode(&newUrl)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(`{"Error": "400: Failed to encode json"}`)
		return
	}
	_, error := url.ParseRequestURI(newUrl.Url)
	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"Error": "406: Not an Acceptable URL"}`)
		return
	}

	//implimented mutexes for no creation overlaps
	mutex.Lock()

	newString := RandomString(8)
	urlMap[newString] = newUrl.Url

	mutex.Unlock()

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, `{"short_url": "http://127.0.0.1:5555/%s"}`, newString)
	log.Println(urlMap)
	return
}
