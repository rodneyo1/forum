package utils

import "net/http"

// Checks if user is loged in
func IsLoggedIn(r *http.Request) bool {
	session, _ := r.Cookie("session_id")
	return session != nil
}
