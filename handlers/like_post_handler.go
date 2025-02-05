package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"forum/database"
)

func LikePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()
	postIDStr := r.FormValue("post_id")
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		fmt.Println("Invalid PostID")
	} else {
		userID := 1 // Replace with actual logged-in user ID

		err := database.Like(userID, &postID, nil)
		if err != nil {
			http.Error(w, "Failed to like post", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
