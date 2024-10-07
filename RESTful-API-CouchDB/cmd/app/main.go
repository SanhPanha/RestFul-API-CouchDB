package main

import (
	"RESTful-API/routes"
	"log"
	"net/http"
)

func main() {
    router := routes.SetupRoutes()

    log.Println("Server starting on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", router))
}
