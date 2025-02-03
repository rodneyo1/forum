package database

import "forum/models"

func GetAllPosts() ([]models.Post, error) {
	query := `SELECT uuid, title, content, media, user_id, created_at FROM posts`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.UUID, &post.Title, &post.Content, &post.Media, &post.UserID, &post.CreatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
