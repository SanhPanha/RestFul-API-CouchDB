package routes

import (
	"RESTful-API/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
    router := mux.NewRouter()

    router.HandleFunc("/users", controllers.CreateUser).Methods("POST")
    router.HandleFunc("/accounts", controllers.CreateAccount).Methods("POST")

    // Other routes for CRUD operations

    return router
}
