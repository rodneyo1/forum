package main

import (
	"log"
	"net/http"

	"forum/handlers"
)

func main() {
	// Candle hundler functions
	http.HandleFunc("/static/", handlers.StaticHandler)
	// http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/success", handlers.SuccessHandler)
	http.HandleFunc("/register", handlers.RegistrationHandler)

	log.Println("Server started on port 8080")

	// Start the server, handle emerging errors
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("Failed to start server: ", err)
		return
	}

}
