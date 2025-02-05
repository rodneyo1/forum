package database

import (
	"fmt"

	"forum/models"
)

// Fetch comments for a post by postUUID
func GetPostsComments(postUUID string) ([]models.CommentWithCreator, error) {
	// SQL query to fetch comments for a specific post
	query := `SELECT c.uuid, c.content, c.post_id, u.username, c.created_at 
	FROM comments c
	INNER JOIN users u ON u.id = c.user_id
	WHERE post_id = ?`

	// Execute the query with the provided postUUID
	rows, err := db.Query(query, postUUID)
	if err != nil {
		// Log and return the error if the query fails
		fmt.Println("Error executing query: ", err)
		return nil, err
	}
	defer rows.Close()

	// Initialize a slice to hold the comments
	var comments []models.CommentWithCreator

	// Loop through the rows and scan each one into the Comment struct
	for rows.Next() {
		var comment models.CommentWithCreator
		err := rows.Scan(&comment.UUID, &comment.Content, &comment.PostID, &comment.Creator, &comment.CreatedAt)
		if err != nil {
			// Log and return error if scanning fails
			fmt.Println("Error scanning row: ", err)
			return nil, err
		}

		// Append each comment to the comments slice
		comments = append(comments, comment)
	}

	// Return the list of comments
	return comments, nil
}
