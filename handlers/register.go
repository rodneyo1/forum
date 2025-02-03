package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"forum/utils"
)

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the method is allowed (e.g., POST)
	if MethodCheck(w, r, http.MethodGet) {
		// Render the registration form
		tmpl, err := template.ParseFiles("web/templates/register.html")
		if err != nil {
			http.Error(w, "Unable to load registration page", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
		return
	}

	if MethodCheck(w, r, http.MethodPost) {
		// Parse form values and validate user input
		user, err := utils.PaerseUserForm(r)
		if err != nil {
			http.Error(w, err.Issue, err.Code)
			return
		}
		utils.Passwordhash(user)
	}
	tmpl, err := template.ParseFiles("web/templates/register.html")
	if err != nil {
		http.Error(w, "Unable to parse registration template", http.StatusInternalServerError)
		return
	}

	// Render registration form when method is GET
	if r.Method == "GET" {
		err := tmpl.Execute(w, nil)
		if err != nil {
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
		ParseAlertMessage(w, tmpl, "Invalid email format")
		return
	}

	// Check if email or username is taken
	existingUser, _ := database.GetUserByEmailOrUsername(r.FormValue("email"), r.FormValue("username"))
	if existingUser.Username != "" {
		ParseAlertMessage(w, tmpl, fmt.Sprintf("%s taken!", r.FormValue("email")))
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
		ParseAlertMessage(w, tmpl, "Passwords do not match")
		return
	}

	user.Password = password // set password

	// Create new user in the database
	_, err = database.CreateUser(user.Username, user.Email, user.Password)
	if err != nil {
		ParseAlertMessage(w, tmpl, "Error creating user")
		return
	}

	// Redirect user to login page
	if w.Header().Get("Content-Type") == "" {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}

func ValidEmail(email string) bool {
	return strings.Contains(email, "@") || strings.HasPrefix(email, ".com")
}
