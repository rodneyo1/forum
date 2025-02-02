package handlers

import (
	"log"
	"net/http"
	"strings"
	"text/template"
	"time"

	"forum/database"
	"forum/models"
)

// LoginHandler handles user login and session creation.
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// block to prevent login in when already logged in
	{
		cookie, err := r.Cookie("session_id")
		if err != nil {
			// No session ID cookie found, redirect to login
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		sessionID := cookie.Value

		// Check if the session exists and is valid
		_, err = database.GetSession(sessionID)
		if err == nil {
			// Session found or already logged in, redirect to home
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	} // end block to prevent login in when already logged in

	var user models.User

	// Check if the method is GET, if so render the login template
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("web/templates/login.html")
		if err != nil {
			http.Error(w, "Failed to load Login template", http.StatusInternalServerError)
			return
		}
		if err := tmpl.Execute(w, nil); err != nil {
			InternalServerErrorHandler(w)
			return
		}
		return
	}

	// Restrict to POST for login submission
	if r.Method != "POST" {
		BadRequestHandler(w)
		log.Println("LoginHandler ERROR: Bad request method")
		return
	}

	// Parse the form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	// Populate user credentials
	// Determine whether input is a valid email
	if IsValidEmail(r.FormValue("email_username")) {
		user.Email = r.FormValue("email_username")
	} else {
		user.Username = r.FormValue("email_username")
	}

	user.Password = r.FormValue("password") // Populate password field
	log.Println("Username: ", user.Username)
	log.Println("Email: ", user.Email)
	log.Println("Password: ", user.Password)

	// Attempt to log in the user
	sessionID, err := database.LoginUser(user.Username, user.Email, user.Password)
	if err != nil {
		// Login failed, render the login template with an error message
		tmpl, err := template.ParseFiles("web/templates/login.html")
		if err != nil {
			http.Error(w, "Failed to load Login template", http.StatusInternalServerError)
			return
		}

		// Pass the error message to the template
		data := struct {
			Error string
		}{
			Error: "Invalid username/email or password",
		}

		if err := tmpl.Execute(w, data); err != nil {
			InternalServerErrorHandler(w)
			return
		}
		return
	}

	// Login successful, set the session ID as a cookie
	cookie := http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour), // Session expires in 24 hours
		HttpOnly: true,                           // Prevent client-side script access
		Secure:   true,                           // Ensure cookie is only sent over HTTPS
		SameSite: http.SameSiteStrictMode,        // Prevent CSRF attacks
	}
	http.SetCookie(w, &cookie)

	// Redirect the user to the home page or a protected route
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// package handlers

// import (
// 	"forum/database"
// 	"forum/models"
// 	"html/template"
// 	"log"
// 	"net/http"
// 	"strings"
// )

// func LoginHandler(w http.ResponseWriter, r *http.Request) {
// 	var user models.User

// 	// Restrict requests to "GET" or "POST"
// 	if !(r.Method == "POST" || r.Method == "GET") {
// 		BadRequestHandler(w)
// 		log.Println("LoginHandler ERROR: Bad request")
// 		return
// 	}

// 	// Render the login template with the request object as data
// 	tmpl, err := template.ParseFiles("web/templates/login.html")
// 	if err != nil {
// 		http.Error(w, "Failed to load Login template", http.StatusInternalServerError)
// 		return
// 	}

// 	// Populate user credentials
// 	// Determine whether input is a valid email
// 	if IsValidEmail(r.FormValue("email_username")) {
// 		user.Email = r.FormValue("email_username")
// 	} else {
// 		user.Username = r.FormValue("username")
// 	}

// 	user.Password = r.FormValue("password") // Populate password field
// 	log.Println("Username: ", user.Username)
// 	log.Println("Email: ", user.Email)
// 	log.Println("Password: ", user.Password)

// 	database.LoginUser(user.Username, user.Password)

// 	if err := tmpl.Execute(w, r); err != nil {
// 		InternalServerErrorHandler(w)
// 		return
// 	}
// }

func IsValidEmail(email string) bool {
	return strings.Contains(email, "@") && strings.HasSuffix(email, ".com")
}
