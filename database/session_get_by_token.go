package database

import (
	"net/http"
	"time"

	"forum/models"
)

// GetSession retrieves a session by its session ID.
func GetSession(sessionID string) (models.Session, error) {
	query := `SELECT session_id, user_id, expiry FROM sessions WHERE session_id = ?`
	var session models.Session
	err := db.QueryRow(query, sessionID).Scan(&session.SessionID, &session.UserID, &session.Expiry)
	if err != nil {
		return models.Session{}, err
	}
	return session, nil
}

// Checks if a user is logged in
func IsLoggedIn(r *http.Request) bool {
	sessionCookie, err := r.Cookie("session_id")
	if err != nil {
		return false // no sessions found
	}

	// Fetch session from database
	session, err := GetSession(sessionCookie.Value)
	if err != nil {
		return false // session not found or invalid
	}

	if session.Expiry.Before(time.Now()) {
		return false // session expired
	}

	return true
}
