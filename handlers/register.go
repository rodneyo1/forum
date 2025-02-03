package handlers

import (
	"html/template"
	"log"
	"net/http"
	"regexp"

	"forum/database"
	"forum/models"
	"forum/utils"
)

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	// Handle GET request: render registration page
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("web/templates/register.html")
		if err != nil {
			http.Error(w, "Unable to load registration page", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
		return
	}

	// Only allow POST method for registration
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse user form data
	user, err := utils.ParseUserForm(r)
	if err != nil {
		http.Error(w, err.Issue, err.Code)
		return
	}

	// Validate email format
	if !ValidEmail(user.Email) {
		renderErrorMessage(w, "Invalid email format")
		return
	}

	// Check if email or username is already taken
	existingUser, _ := database.GetUserByEmailOrUsername(user.Email, user.Username)
	if existingUser.Username == user.Username {
		renderErrorMessage(w, "Username already taken")
		return
	}
	if existingUser.Email == user.Email {
		renderErrorMessage(w, "Email already registered")
		return
	}

	// Check password strength
	if err := utils.PasswordStrength(user.Password); err != nil {
		renderErrorMessage(w, "Password too weak")
		return
	}

	// Hash the password before storing in database
	utils.Passwordhash(user) 

	// Store the user in the database
	err = database.CreateUser(user.Username, user.Email, user.Password)
	if err != nil {
		renderErrorMessage(w, "Error creating user")
		return
	}

	// Redirect user to login page
	http.Redirect(w, r, "/login", http.StatusFound)
}

// Validate email format
func ValidEmail(email string) bool {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	return regexp.MustCompile(regex).MatchString(email)
}

// Helper function to render error messages
func renderErrorMessage(w http.ResponseWriter, message string) {
	tmpl, err := template.ParseFiles("web/templates/register.html")
	if err != nil {
		http.Error(w, "Unable to load registration page", http.StatusInternalServerError)
		return
	}
	ParseAlertMessage(w, tmpl, message)
}
