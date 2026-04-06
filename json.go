package main

import (
	"encoding/json"
	"log"
	"net/http"
)



func ResponseWithJson(w http.ResponseWriter, code int, payload interface{}){
	data, err := json.Marshal(payload)
	if err !=nil {
		log.Printf("failed to marshel data: %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}