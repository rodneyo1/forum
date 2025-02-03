package database

import "forum/models"

func GetAllPosts() ([]models.Post, error) {
	return []models.Post{}, nil
}

func Dislike(int, *string, any) error {
	return nil
}

func Like(int, *string, any) error {
	return nil
}
