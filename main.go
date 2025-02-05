package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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
	http.HandleFunc("/home", handlers.ForumHandler)
	http.HandleFunc("/home", handlers.ForumHandler)
	http.Handle("/posts/create", middleware.AuthMiddleware(http.HandlerFunc(postHandlers.PostCreate)))

	// Inform user initialization of server
	log.Println("Server started on port 8080")

	// Start the server, handle emerging errors
	err := http.ListenAndServe(":8080", nil)
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
