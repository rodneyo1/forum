package database

import (
	"forum/models"
)

// func GetPostByUUID(postID string) (models.Post, error) {
// 	query := `SELECT p.uuid, p.title, p.content, p.media, u.username, p.user_id, p.created_at
// 	FROM posts p
// 	INNER JOIN users u ON p.user_id = u.id
// 	INNER JOIN post_categories pc ON p.uuid = pc.post_id
// 	INNER JOIN categories c
	
// 	FROM posts WHERE id = ?`
// 	var post models.PostWithCategories
// 	err := db.QueryRow(query, postID).Scan(&post.UUID, &post.Title, &post.Content, &post.Media, &post.Username, &post.CreatedAt)
// 	if err != nil {
// 		return models.Post{}, err
// 	}
// 	return post, nil
// }

func GetPostByUUID(postID string) (models.PostWithCategories, error) {
	var post models.PostWithCategories
	return post, nil
}