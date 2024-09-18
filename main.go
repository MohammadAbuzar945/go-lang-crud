package main

import (
	"bookManagement/config"
	"bookManagement/routes"
	"log"
	"net/http"
)

func main() {
	config.InitializaDB()

	config.InitilizeS3()
	r := routes.InitializeRoutes()

	log.Println("Server started on: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
