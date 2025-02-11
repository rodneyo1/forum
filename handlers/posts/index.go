package posts

import (
	"html/template"
	"log"
	"net/http"

	"forum/database"
	"forum/models"
)

func Index(w http.ResponseWriter, r *http.Request) {
	loggedIn := false
	session, lIn := database.IsLoggedIn(r)
	if lIn {
		loggedIn = true
	}

	// Retrieve user data
	userData, err := database.GetUserbySessionID(session.SessionID)

	if err != nil {
		log.Printf("Error getting user: %v\n", err) // Add error logging
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Fetch posts from the database
	posts, err := database.GetAllPosts()
	if err != nil {
		http.Error(w, "Failed to load posts", http.StatusInternalServerError)
		return
	}
	categories, err := database.FetchCategories()
	if err != nil {
		http.Error(w, "Failed to fetch categories", http.StatusInternalServerError)
		return
	}
	// Load the HTML template
	tmpl := template.Must(template.ParseFiles("web/templates/index.html"))
	// Execute the template with the posts data
	data := struct {
		Posts      []models.PostWithUsername
		Categories []models.Category
		IsLogged   bool
		ProfPic    string
	}{
		Posts:      posts,
		Categories: categories,
		IsLogged:   loggedIn,
		ProfPic:    userData.Image,
	}

	tmpl.Execute(w, data)
}
