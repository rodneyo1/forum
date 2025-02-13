package posts

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"forum/database"
	"forum/models"
)

type TemplateData struct {
	IsLogged bool
	Posts    []models.PostWithCategories
	ProfPic  string
}

func ShowLikedPosts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	userID, _, err := database.GetUserData(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Retrieve user information
	user, _ := database.GetUserbyID(userID)

	userIDStr := strconv.Itoa(userID)

	likedPosts, err := database.GetLikedPostsByUser(userIDStr)
	if err != nil {
		http.Error(w, "Failed to retrieve liked posts", http.StatusInternalServerError)
		return
	}

	data := TemplateData{
		IsLogged: true,
		Posts:    likedPosts,
		ProfPic:  user.Image,
	}

	tmpl, err := template.ParseFiles("web/templates/liked-posts.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	// Render the liked posts page
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("Error executing template:", err)
		return
	}
}
