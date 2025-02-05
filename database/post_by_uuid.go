package database

import (
	"fmt"
	"strings"

	"forum/models"
)

func GetPostByUUID(postID string) (models.PostWithCategories, error) {
	query := `
		SELECT 
			p.uuid, 
			p.title, 
			p.content, 
			p.media, 
			u.username, 
			p.user_id, 
			p.created_at,
			GROUP_CONCAT(c.name) AS category_names, -- Aggregate category names
			COUNT(DISTINCT l.id) AS likes_count,    -- Count distinct likes
			COUNT(DISTINCT dl.id) AS dislikes_count -- Count distinct dislikes
		FROM posts p
		INNER JOIN users u ON p.user_id = u.id
		LEFT JOIN post_categories pc ON p.uuid = pc.post_id
		LEFT JOIN categories c ON pc.category_id = c.id
		LEFT JOIN likes l ON p.uuid = l.post_id
		LEFT JOIN dislikes dl ON p.uuid = dl.post_id
		WHERE p.uuid = ?
		GROUP BY p.uuid, p.title, p.content, p.media, u.username, p.user_id, p.created_at
	`

	// Execute the query
	row := db.QueryRow(query, postID)

	var post models.PostWithCategories
	var categoryNames string // Temporary variable to hold aggregated category names

	// Scan the result into the post struct
	err := row.Scan(
		&post.UUID,
		&post.Title,
		&post.Content,
		&post.Media,
		&post.Username,
		&post.UserID,
		&post.CreatedAt,
		&categoryNames, // Scan aggregated category names
		&post.LikesCount,
		&post.DislikesCount,
	)
	if err != nil {
		return models.PostWithCategories{}, fmt.Errorf("failed to fetch post: %v", err)
	}

	// Split the aggregated category names into a slice
	if categoryNames != "" {
		post.Categories = strings.Split(categoryNames, ",")
	} else {
		post.Categories = []string{} // Ensure Categories is not nil
	}

	return post, nil
}
