package database

import (
	"forum/models"
)

func GetUserByEmailOrUsername(email, username string) (models.User, error) {
	query := `SELECT id, username, email, password FROM users WHERE email = ? OR username = ?`
	var user models.User
	err := db.QueryRow(query, email, username).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
