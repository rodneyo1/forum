package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"forum/database"
	"forum/handlers"
	postHandlers "forum/handlers/posts"
)

func init() {
	err := database.Init("storage/forum.db")
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func main() {
	port := ":8080"
	portStr := os.Getenv("PORT")

	if _, e := strconv.Atoi(portStr); e == nil {
		port = ":" + portStr
	}
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
	// RESTORE // http.Handle("/posts/create", middleware.AuthMiddleware(http.HandlerFunc(postHandlers.PostCreate)))
	http.HandleFunc("/posts/create", postHandlers.PostCreate)
	http.HandleFunc("/posts/display", postHandlers.PostDisplay)
	// User Profile routes
	http.HandleFunc("GET /profile", handlers.ViewUserProfile)
	// http.HandleFunc("GET /user/update", middleware.AuthMiddleware(http.HandlerFunc(handlers.UpdateUserProfile))) // Protected

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		user, _ := database.GetUserByEmailOrUsername("toni", "toni")
		fmt.Println("User: ", user)
	})

	http.HandleFunc("/posts/like", handlers.LikePostHandler)
	http.HandleFunc("/posts/dislike", handlers.DislikePostHandler)
	http.HandleFunc("/comments/like", handlers.LikeCommentHandler)
	http.HandleFunc("/comments/dislike", handlers.DislikeCommentHandler)
	http.HandleFunc("/comment", Comment)

	// Inform user initialization of server
	log.Printf("Server runing on http://localhost:%s\n", portStr)

	// Start the server, handle emerging errors
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Println("Failed to start server: ", err)
		return
	}
}

func Comment(w http.ResponseWriter, r *http.Request) {
	// take the contents from the form
	// call the database function to insert a comment
	// redirect to /posts/display?pid={{.UUID}}

	// Only allow POST requests for submitting a comment
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse form data (assuming the form contains a comment and post UUID)
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	// Retrieve the comment text and post UUID from the form
	commentText := r.FormValue("comment")
	postUUID := r.FormValue("postUUID")
	// userID := r.FormValue("userID") // Assuming userID is passed in the form
	userID := 1

	// fmt.Println("TEXT: ", commentText)

	// Ensure the comment is not empty
	if commentText == "" {
		http.Error(w, "Comment cannot be empty", http.StatusBadRequest)
		return
	}

	// // Validate the UUID format (basic validation, adjust according to your needs)
	// if !utils.IsValidUUID(postUUID) {
	// 	http.Error(w, "Invalid post UUID", http.StatusBadRequest)
	// 	return
	// }

	// Assuming you convert the userID from string to int
	// userIDInt, err := strconv.Atoi(userID)
	// if err != nil {
	// 	http.Error(w, "Invalid user ID", http.StatusBadRequest)
	// 	return
	// }
	// Call the CreateComment function to insert the comment into the database
	_, err = database.CreateComment(userID, postUUID, commentText)
	if err != nil {
		http.Error(w, "Error inserting comment into database", http.StatusInternalServerError)
		return
	}

	// Redirect to the post's display page with the post UUID
	http.Redirect(w, r, fmt.Sprintf("/posts/display?pid=%s", postUUID), http.StatusSeeOther)
}
