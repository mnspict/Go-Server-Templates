package handlers

import (
	"fmt"
	"net/http"

	"main.go/internal/services"
)

// ExHandler defines the handler for the endpoint "GET/example".
// Handler reads or writes to the request context. 
// Any core logic should be forwarded to the service layer.
func ExHandler(w http.ResponseWriter, r *http.Request) {

	services.ExService()

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to the Go HTTP Server built with Gorilla/Mux!")

}