package main

import (
	"encoding/json"
	"net/http"

	
	// "github.com/MokhtarOmar16/rssagg-GO/internal/database"
)

func (apiCfg *apiconfig)  CreateUserHandler(w http.ResponseWriter, r *http.Request){
	params := struct{
		Name string `json:"name"`
	}{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
        ResponseWithError(w, 400, "Invalid request payload")
        return
    }
	user , err := apiCfg.DB.CreateUser(r.Context(), params.Name)
	if err != nil {
        ResponseWithError(w, 400, "error while creating user")
        return
    }
	ResponseWithJson(w, 201, user)
}




