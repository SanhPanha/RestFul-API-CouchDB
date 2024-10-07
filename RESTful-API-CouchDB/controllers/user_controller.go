package controllers

import (
	"RESTful-API/pkg/services"
	"encoding/json"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
    var user services.User
    json.NewDecoder(r.Body).Decode(&user)

    err := services.CreateUser(user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
    // Logic for getting a user
}

// Other CRUD and bulk operation handlers for users...
