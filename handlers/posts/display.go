package posts

import (
	"html/template"
	"log"
	"net/http"

	"forum/database"
	"forum/models"
)

func PostDisplay(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./web/templates/post_display.html")
	if err != nil {
		http.Error(w, "Failed to load post template", http.StatusInternalServerError)
		return
	}
	postID := r.URL.Query().Get("pid")
	// fmt.Println("SINGLE PID: ", postID)

	postData, err := database.GetPostByUUID(postID)
	if err != nil {
		log.Println("Error getting post data: ", err)
		return
	}

	// Infuse data to be executed with inquiry if user is logged in
	data := struct {
		PostData   models.PostWithCategories
		IsLoggedIn bool
	}{
		PostData:   postData,
		IsLoggedIn: database.IsLoggedIn(r),
	}

	// fmt.Println("POST: ", PostData)

	if err := tmpl.Execute(w, data); err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
