package handlers

import (
	"net/http"

	"forum/database"
)

func LikePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()
	postID := r.FormValue("post_id")
	userID := 1 // Replace with actual logged-in user ID

	err := database.Like(userID, postID, postID)
	if err != nil {
		http.Error(w, "Failed to like post", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
