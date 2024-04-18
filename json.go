package main

import (
	"encoding/json"
	"log"
	"net/http"
)

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

func respondWithJSON(w http.ResponseWriter, code int, paylode interface{}) {
	data, err := json.Marshal(paylode)
	if err != nil {
		log.Printf("failed to marshal JSON response : %v", paylode)
		w.WriteHeader(500)
		return

	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
