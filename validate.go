package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func (cfg *apiConfig) handlerValidate(w http.ResponseWriter, r *http.Request) {
	type skeleton struct {
		Body string `json:"body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := skeleton{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("Error decoding parameters: %s", err)
		w.WriteHeader(500)
		return
	}
}

func (cfg *apiConfig) handlerEncoding(w http.ResponseWriter, r *http.Request) {
	
}
