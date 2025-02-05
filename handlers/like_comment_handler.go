package handlers

import (
	"net/http"

	"forum/database"
)

func LikeCommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()
	commentID := r.FormValue("comment-id")

	userID := 1 // Replace with actual logged-in user ID

	err := database.LikeComment(userID, commentID)
	if err != nil {
		http.Error(w, "Failed to like comment", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
