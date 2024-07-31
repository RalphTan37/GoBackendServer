package main

import (
	"encoding/json"
	"log"
	"net/http"
)

//JSON rest api - all the request bodies coming in and back will have a JSON format

// responding w/ error messages
func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5XX error:", msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, code, errResponse{
		Error: msg,
	})
}

// helper function that makes it easier to send JSON responses
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	//marshal the payload into a JSON object/string
	data, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500) //"something went wrong on my end" - internal server error
		return
	}

	w.Header().Add("Content-Type", "application/json") //responding with JSON data
	w.WriteHeader(code)                                //write status code - "everything went well"
	w.Write(data)                                      //writes the response body
}
