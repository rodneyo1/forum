package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"forum/database"
)

func LikeCommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()
	commentIDStr := r.FormValue("comment_id")
	commentID, err := strconv.Atoi(commentIDStr) // Convert to int
	if err != nil {
		fmt.Println("Invalid comment_id")
	} else {
		userID := 1 // Replace with actual logged-in user ID

		err := database.Like(userID, nil, &commentID)
		if err != nil {
			http.Error(w, "Failed to like comment", http.StatusInternalServerError)
			return
		}
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
