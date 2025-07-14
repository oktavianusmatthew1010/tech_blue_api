// main.go
// @title Caregiver Scheduling API
// @version 1.0
// @description API for Tech Blue API
// @contact.name Oktavianus
// @contact.email oktavianus.programmer@gmail.com
// @license.name MIT
// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	routes "github.com/oktavianusmatthew1010/tech_blue_api/routers"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()

	routes.SetupRoutes(router)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://tech-blue-fe.vercel.app"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)
	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
