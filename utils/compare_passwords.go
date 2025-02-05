package utils

import "golang.org/x/crypto/bcrypt"

// Uses bcrypt to compare hashed and non-hashed passwords
func MatchPasswords(hashedPassword, inputPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
	return err == nil
}
