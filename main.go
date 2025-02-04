package main

import (
	"log"
	"net/http"
	"os"

	"forum/database"
	"forum/handlers"
	middleware "forum/handlers/middlewares"
	postHandlers "forum/handlers/posts"
)

func init() {
	err := database.Init("storage/forum.db")
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func main() {
	defer database.Close()

	database.CreateUser("toni", "toni@mail.com", "@antony222")

	// Restrict arguments parsed
	if len(os.Args) != 1 {
		log.Println("Too many arguments")
		log.Println("Usage: go run main.go")
		return
	}

	// Candle hundler functions
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/static/", handlers.StaticHandler)
	http.HandleFunc("/success", handlers.SuccessHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/forgot-password", handlers.ForgotPasswordHandler)
	http.HandleFunc("/register", handlers.RegistrationHandler)
	http.Handle("/posts/create", middleware.AuthMiddleware(http.HandlerFunc(postHandlers.PostCreate)))
	// http.Handle("/posts/", middleware.AuthMiddleware(http.HandlerFunc(postHandlers.DisplaySinglePost)))
	http.HandleFunc("/posts/display", postHandlers.PostDisplay)

	// Inform user initialization of server
	log.Println("Server running on http://localhost:8080")

	// Start the server, handle emerging errors
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("Failed to start server: ", err)
		return
	}
}
