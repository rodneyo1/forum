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

// func Authorized(r *http.Request) error {
// 	// Check if session cookie exists
// 	cookie, err := r.Cookie("session_id")
// 	if err != nil {
// 		if err == http.ErrNoCookie {
// 			// no cookie found
// 			return errors.New("no cookie found")
// 		}
// 		// if there's an actual error while reading the cookie
// 		return errors.New("internal server error")
// 	}

// 	// Session cookie exists, check if user is logged in
// 	session, err := GetSession(cookie.Value)
// 	if err != nil {
// 		// Session not found or invalid
// 		return errors.New("session not found or invalid")
// 	}

// 	if session.Expiry.Before(time.Now()) {
// 		// session expired
// 		return errors.New("session expired")
// 	}

// 	// session is valid, user is logged in, handle cookie
// 	hasCookie, _, err := HasCookie(r)
// 	if err != nil {
// 		// there's an error while checking the cookie
// 		fmt.Printf("Error checking session cookie", err)
// 		return errors.New("error checking session cookie")
// 	}

// 	if hasCookie {
// 		// success
// 		return nil
// 	}

// 	// No cookie found, user not logged in
// 	return errors.New("an error occured while retrieving the cookie")
// }
