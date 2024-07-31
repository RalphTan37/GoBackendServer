package main

import "net/http"

//define an HTTP handler
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{}) //passes the HTTP response writer, respond w/ 200 status code, empty JSON object
}
