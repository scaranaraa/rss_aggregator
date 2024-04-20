package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func resopondWithError(w http.ResponseWriter, code int, msg string){

	if code > 499 {
		log.Println("Error 5XX", msg)
	}

	type ErrorResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, code, ErrorResponse{
		Error: msg,
	})

}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}){

	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v",payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(code)
	w.Write(dat)
}