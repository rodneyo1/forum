package database

import (
	"forum/models"
)

// fetches all posts from the database with the creator's names
func GetAllPosts() ([]models.PostWithUsername, error) {
	query := `SELECT p.uuid, p.title, p.content, p.media, p.user_id, p.created_at, u.username FROM posts p JOIN users u ON u.id = p.user_id`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.PostWithUsername
	for rows.Next() {
		var post models.PostWithUsername
		err := rows.Scan(&post.UUID, &post.Title, &post.Content, &post.Media, &post.UserID, &post.CreatedAt, &post.Creator)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}
