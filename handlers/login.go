package handlers

import (
	"forum/database"
	"forum/models"
	"forum/utils"
	"html/template"
	"log"
	"net/http"
	"time"
)

// LoginHandler handles user login and session creation, as well as preventing login when already logged in.
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
	// Build path to login.html
	templatePath, err := GetTemplatePath("login.html")
	if err != nil {
		InternalServerErrorHandler(w)
		log.Println("Could not find template file: ", err)
		return
	}

	// If the method is GET, if serve blank login form
	if r.Method == "GET" {
		// Skip login for users who are already loged in
		{
			hasCookie := HasCookie(r, w)
			if hasCookie {
				http.Redirect(w, r, "/", http.StatusFound)
				return
			}
		}

		// Render login form
		tmpl, err := template.ParseFiles(templatePath)
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

	// Catch non-Get and non-POST requests
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
	if utils.IsValidEmail(r.FormValue("email_username")) {
		user.Email = r.FormValue("email_username")
	} else {
		user.Username = r.FormValue("email_username")
	}
	// Extract form data
	user.Password = r.FormValue("password") // Populate password field
	emailUsername := r.FormValue("email_username")
	password := r.FormValue("password")

	// Validate credentials
	if !database.VerifyUser(emailUsername, password) {
		http.Redirect(w, r, "/login", http.StatusOK)
		return
	}

	// Attempt to log in the user
	sessionID, err := database.LoginUser(user.Username, user.Email, user.Password)
	if err != nil {
		// Login failed, render the login template with an error message
		templatePath, err := GetTemplatePath("login.html")
		if err != nil {
			InternalServerErrorHandler(w)
			log.Println("Could not find template file: ", err)
			return
		}
		// Render error message
		tmpl, err := template.ParseFiles(templatePath)
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

func HasCookie(r *http.Request, w http.ResponseWriter) bool {
	cookie, err := r.Cookie("session_id")
	log.Printf(" Cookie : %#v, err: %#v", cookie, err)
	if err != nil {
		// No session ID cookie found, redirect to login
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return true
	}

	sessionID := cookie.Value

	// Check if the session exists and is valid
	_, err = database.GetSession(sessionID)
	if err == nil {
		// Session found or already logged in, redirect to home
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return true
	}

	return false
}

// ParseAlertMessage is used for displaying alert messages in templates.
func ParseAlertMessage(w http.ResponseWriter, tmpl *template.Template, message string) {
	// Define template path and error message
	alert := map[string]string{"ErrorMessage": message}
	// Execute the page
	err := tmpl.Execute(w, alert)
	if err != nil {
		InternalServerErrorHandler(w)
		log.Printf("Could not execute template %v", err)
		return
	}
}

func SuccessHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		BadRequestHandler(w)
		return
	}
	http.Error(w, "Loged in1!", http.StatusInternalServerError)
}
