package auth

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"forum/database"
	"forum/models"
	"forum/utils"
)

// ViewUserProfile handler
func ViewProfile(w http.ResponseWriter, r *http.Request) {
	session, ok := database.SessionFromContext(r)
	if !ok {
		log.Printf("Error getting user: %v\n", session) // Add error logging
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	// fmt.Printf("UserData retrieved: %+v\n", session) // Add debug logging

	user, err := database.GetUserbyID(session.UserID)
	if err != nil {
		http.Error(w, "Error getting user profile", http.StatusInternalServerError)
	}

	var posts []models.Post
	posts, err = database.GetPostsByUserID(session.UserID)
	if err != nil {
		posts = nil
	}

	// Render the template with data
	path, err := utils.GetTemplatePath("profile.html")
	if err != nil {
		fmt.Println("Error getting template path")
	}

	tmpl, err := template.ParseFiles(path)
	if err != nil {
		http.Error(w, "Failed to load profile template", http.StatusInternalServerError)
		return
	}

	data := struct {
		User  models.User
		Posts []models.Post
	}{
		User:  user,
		Posts: posts,
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// func UpdateUserProfile(){
// 	// Update user profile
// }
