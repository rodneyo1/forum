package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"forum/database"
	"forum/models"
)

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	// Construct absolute path to register.html
	tmplPath, err := GetTemplatePath("register.html")
	if err != nil {
		InternalServerErrorHandler(w)
		log.Println("Could not find template file: ", err)
		return
	}

	// Render html template
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Unable to parse registration template", http.StatusInternalServerError)
		return
	}

	// Render registration form when method is GET
	if r.Method == "GET" {
		if err := tmpl.Execute(w, r); err != nil {
			InternalServerErrorHandler(w)
			log.Println("Could not render registration template: ", err)
			return
		}
		return
	}

	// Handle non-GET and non-POST requests
	if r.Method != "POST" {
		BadRequestHandler(w)
		log.Println("RegistrationHandler ERROR: Bad request")
		return
	}

	// Parse form data
	err = r.ParseForm()
	if err != nil {
		BadRequestHandler(w)
		log.Println("Invalid form submission", http.StatusBadRequest)
		return
	}

	// Validate email format
	if !ValidEmail(r.FormValue("email")) {
		ParseAlertMessage(w, tmpl, tmplPath, "Invalid email format")
		return
	}

	// Check if email or username is taken
	existingUser, _ := database.GetUserByEmailOrUsername(r.FormValue("email"), r.FormValue("username"))
	if existingUser.Username != "" {
		ParseAlertMessage(w, tmpl, tmplPath, fmt.Sprintf("%s taken!", r.FormValue("email")))
		return
	}

	// Extract form data
	user.Email = r.FormValue("email")
	user.Username = r.FormValue("username")
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirm_password")
	user.Bio = r.FormValue("bio")

	// Validate passwords
	if password != confirmPassword {
		ParseAlertMessage(w, tmpl, tmplPath, "Passwords do not match")
		return
	}

	user.Password = password // set password

	// Create new user in the database
	_, err = database.CreateUser(user.Username, user.Email, user.Password)
	if err != nil {
		ParseAlertMessage(w, tmpl, tmplPath, "Error creating user")
		return
	}

	// Redirect user to login page
	ParseAlertMessage(w, tmpl, tmplPath, fmt.Sprintf("%v, you created a new account", user.Username))
	http.Redirect(w, r, "/login", http.StatusFound)
}

func ValidEmail(email string) bool {
	return strings.Contains(email, "@") || strings.HasPrefix(email, ".com")
}
