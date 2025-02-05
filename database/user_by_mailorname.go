package database

import (
	"fmt"

	"forum/models"
	"forum/utils"
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

func VerifyUser(email, password string) (bool, error) {
	// User email or username to get users' full credentials
	user, err := GetUserByEmailOrUsername(email, email)
	if err != nil {
		return false, fmt.Errorf("user does not exist: %v", err)
	}

	// Compare provided password with the stored hashed password
	return utils.MatchPasswords(user.Password, password)
}
