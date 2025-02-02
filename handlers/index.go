package handlers

import (
	"encoding/json"
	"forum/database"
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("web/templates/index.html"))

// HomeHandler serves the main forum page with posts
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := database.GetAllPosts() // You need to implement this function in `database` package
	if err != nil {
		http.Error(w, "Unable to retrieve posts", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, posts)
}
func LikePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	userID := 1 // This should be retrieved from session
	postID := r.FormValue("post_id")

	err := database.Like(userID, &postID, nil)
	if err != nil {
		http.Error(w, "Failed to like post", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func DislikePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	userID := 1 // This should be retrieved from session
	postID := r.FormValue("post_id")

	err := database.Dislike(userID, &postID, nil)
	if err != nil {
		http.Error(w, "Failed to dislike post", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// FetchCategoriesHandler handles requests to fetch categories
func FetchCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch categories from the database
	categories, err := database.FetchCategories()
	if err != nil {
		http.Error(w, "Failed to fetch categories", http.StatusInternalServerError)
		return
	}

	// Set response header
	w.Header().Set("Content-Type", "application/json")

	// Encode categories into JSON and send response
	if err := json.NewEncoder(w).Encode(categories); err != nil {
		http.Error(w, "Failed to encode categories", http.StatusInternalServerError)
	}
}
