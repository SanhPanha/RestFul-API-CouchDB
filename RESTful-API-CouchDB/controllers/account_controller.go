package controllers

import (
	"RESTful-API/pkg/services"
	"encoding/json"
	"net/http"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {
    var account services.Account
    json.NewDecoder(r.Body).Decode(&account)

    err := services.CreateAccount(account)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
}

// Other CRUD and bulk operation handlers for accounts...
