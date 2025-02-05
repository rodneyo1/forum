package handlers

import (
	"log"
	"net/http"

	"forum/utils"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Restrict non-POST requests
	if r.Method != "POST" {
		BadRequestHandler(w)
		log.Println("ERROR: LogoutHandler does not allow non-POST requests")
		return
	}

	// Destroy cookies if any, logout session from database
	err := utils.LogOutSession(w, r)
	if err != nil {
		InternalServerErrorHandler(w)
		return
	}

	// Redirect user to home page
	http.Redirect(w, r, "/home", http.StatusFound)
}
