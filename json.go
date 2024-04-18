package main

import (
	"encoding/json"
	"log"
	"net/http"
)

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
