package handlers

import (
	"log"
	"net/http"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Restrict non-POST requests
	if r.Method != "POST" {
		BadRequestHandler(w)
		log.Println("LogoutHandler does not allow non-POST requests")
	}

	// Destroy cookies if any
	{
		hasCookie, err := HasCookie(r)
		if err != nil {
			InternalServerErrorHandler(w)
			log.Println("Error checking session cookie: ", err)
		}

		if hasCookie {
			http.Redirect(w, r, "/home", http.StatusFound)
			return
		}
	}
}
