package handlers

import (
	"net/http"
	"strconv"

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

	postIDStr := r.FormValue("post_id")
	commentIDStr := r.FormValue("comment_id")

	// Replace with actual logged-in user ID
	userID := 1

	// Convert postID and commentID from string to *int
	var postID, commentID *int

	if postIDStr != "" {
		parsedPostID, err := strconv.Atoi(postIDStr)
		if err != nil {
			http.Error(w, "Invalid post ID", http.StatusBadRequest)
			return
		}
		postID = &parsedPostID
	}

	if commentIDStr != "" {
		parsedCommentID, err := strconv.Atoi(commentIDStr)
		if err != nil {
			http.Error(w, "Invalid comment ID", http.StatusBadRequest)
			return
		}
		commentID = &parsedCommentID
	}

	// Call the Dislike function with the database connection
	err = database.Dislike(database.DB, userID, postID, commentID)
	if err != nil {
		http.Error(w, "Failed to dislike post/comment", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
