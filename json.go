package main

import (
	"encoding/json"
	"log"
	"net/http"
)


func ResponseWithError(w http.ResponseWriter, code int, msg string){
	if code > 499 {
		log.Println("Responding with 5xx error: ", msg)
	}
	type ErrResponse struct{
		Error string `json:"error"`
	}
	ResponseWithJson(w, code,	ErrResponse{
		Error: msg,
	} )
}

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