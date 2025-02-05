package handlers

import (
	"net/http"

	"forum/database"
)

func DislikePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	postID := r.FormValue("post_id")
	// commentID := r.FormValue("comment_id")

	// Replace with actual logged-in user ID
	userID := 1

	// Call the Dislike function with the database connection
	err = database.Dislike(userID, postID, postID)
	if err != nil {
		http.Error(w, "Failed to dislike post/comment", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
