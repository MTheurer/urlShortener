package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
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
	mutex := &sync.Mutex{}
	log.Println("starting creatKey server")
	w.Header().Set("Content-Type", "application/json")

	var newKey urlFormat

	fmt.Println("decoding newKey")
	err := json.NewDecoder(r.Body).Decode(&newKey)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(`{"Error": "400: "Failed to encode json}`)
		return
	}

	//adds entry to new struct, compares to Pitchfork struct, outputs appropiate message
	if _, ok := urlMap[newKey.Url]; ok {
		w.WriteHeader(http.StatusConflict)
		fmt.Fprintf(w, `{"Conflict": "409: URL Already Exists"}`)
		return
	}

	//implimented mutexes for no creation overlaps
	mutex.Lock()

	newString := RandomString(8)
	urlMap[newString] = newKey.Url

	mutex.Unlock()

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, `{"short_url": "http://127.0.0.1:5555/%s"}`, newString)
	log.Println(urlMap)
	return
}
