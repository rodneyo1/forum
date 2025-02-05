// package posts

// import (
// 	"html/template"
// 	"net/http"

// 	"forum/database"
// )

// // Handler for serving the form and handling form submission
// func PostDisplay(w http.ResponseWriter, r *http.Request) {
// 	queryParams := r.URL.Query()

// 	postUUID := queryParams.Get("pid")

// 	switch r.Method {
// 	case http.MethodGet:
// 		// Fetch categories from the database
// 		post, err := database.GetPostByUUID(postUUID)
// 		if err != nil {
// 			http.Error(w, "Failed to fetch categories", http.StatusInternalServerError)
// 			return
// 		}

// 		// Serve the form for creating a post
// 		tmpl := template.Must(template.ParseFiles("./web/templates/post_display.html"))
// 		tmpl.Execute(w, post)

// 	default:
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 	}
// }

package posts

import (
	"html/template"
	"log"
	"net/http"

	"forum/database"
)

func PostDisplay(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./web/templates/post_display.html")
	if err != nil {
		http.Error(w, "Failed to load post template", http.StatusInternalServerError)
		return
	}
	postID := r.URL.Query().Get("pid")
	// fmt.Println("SINGLE PID: ", postID)

	PostData, err := database.GetPostByUUID(postID)
	if err != nil {
		log.Println("Error getting post data: ", err)
		return
	}

	// fmt.Println("POST: ", PostData)

	if err := tmpl.Execute(w, PostData); err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
