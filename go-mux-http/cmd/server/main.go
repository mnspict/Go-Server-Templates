package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"main.go/cmd"
	"main.go/internal/handlers"
	"main.go/internal/middlewares"
)


func main() {
    
    // Load all enviornment variables.
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading environment variables: %v\n", err)
    }

    // Initialize database and cache connections.
    if err := cmd.InitDB(); err != nil {
        log.Fatalf("Failed to initialize DB: %v\n", err)
    }

    // Create a new Router.
	router := mux.NewRouter()
    // This implementation uses Gorilla/Mux as the router which doesn't have any in-built middlewares.
    // Attach any middlewares here.
    router.Use(middlewares.Logger) 

    // Load all routes 
    if err := loadRoutes(router); err != nil {
        log.Fatalf("Error loading routes: %v\n", err)
    }

    // Start and listen on server
    port, exists := os.LookupEnv("PORT")
    if !exists || port == "" {
        log.Fatalf("PORT not given in env: %v\n", port)
    }
    fmt.Printf("Server Listening on port %v ... \n", port)
    log.Fatal(http.ListenAndServe(":" + port, router))
}


// loadRoutes registers all the routes to their specific handlers
func loadRoutes(router *mux.Router) error {

    router.HandleFunc("/example", handlers.ExHandler).Methods("GET")
    // Add desired routes here.


    return nil
}