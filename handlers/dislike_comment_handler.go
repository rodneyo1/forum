package handlers

import (
	"net/http"

	"forum/database"
)

func DislikeCommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()
	commentID := r.FormValue("comment_id")
	userID := 1 // Replace with actual logged-in user ID

	err := database.Dislike(userID, nil, &commentID)
	if err != nil {
		http.Error(w, "Failed to dislike comment", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
