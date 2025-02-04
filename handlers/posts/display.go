package posts

import (
	"fmt"
	"html/template"
	"net/http"

	"forum/database"
)

// Handler for serving the form and handling form submission
func PostDisplay(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	postUUID := queryParams.Get("pid")
	fmt.Println("UUID: ", postUUID)

	switch r.Method {
	case http.MethodGet:
		// Fetch categories from the database
		post, err := database.GetPostByUUID(postUUID)
		if err != nil {
			http.Error(w, "Failed to fetch categories", http.StatusInternalServerError)
			return
		}

		// Serve the form for creating a post
		tmpl := template.Must(template.ParseFiles("./web/templates/post_display.html"))
		tmpl.Execute(w, post)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
