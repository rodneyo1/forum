package handlers

import (
	"html/template"
	"log"
	"net/http"
	"regexp"

	"forum/database"
	"forum/utils"
)

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	user, err := utils.ParseUserForm(r)

	if err != nil {
		http.Error(w, err.Issue, err.Code)
		return
	}
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

	//if method is Post updated
	if MethodCheck(w, r, http.MethodPost) {
		//update hashed password
		utils.Passwordhash(user)
	}

	// Handle non-GET and non-POST requests
	if r.Method != "POST" && r.Method != "GET" {
		BadRequestHandler(w)
		log.Println("RegistrationHandler ERROR: Bad request")
		return
	}

	// Validate email format
	if !ValidEmail(r.FormValue("email")) {
		tmpl, err := template.ParseFiles("web/templates/register.html")
		if err != nil {
			http.Error(w, "Unable to load registration page", http.StatusInternalServerError)
			return
		}
		ParseAlertMessage(w, tmpl, "Invalid email format")
		return
	}

	// Check if email or username is taken
	existingUser, _ := database.GetUserByEmailOrUsername(r.FormValue("email"), r.FormValue("username"))
	email := existingUser.Email
	username := existingUser.Username
	if existingUser.Username == username {
		tmpl, err := template.ParseFiles("web/templates/register.html")
		if err != nil {
			http.Error(w, "Unable to load registration page", http.StatusInternalServerError)
			return
		}
		ParseAlertMessage(w, tmpl, "Invalid email format")
		return
	}

	if existingUser.Email == email {
		tmpl, err := template.ParseFiles("web/templates/register.html")
		if err != nil {
			http.Error(w, "Unable to load registration page", http.StatusInternalServerError)
			return
		}
		ParseAlertMessage(w, tmpl, "Invalid email format")
		return
	}

	// Check password strength
	if err := utils.PasswordStrength(user.Password); err != nil {
		tmpl, err := template.ParseFiles("web/templates/register.html")
		if err != nil {
			http.Error(w, "Unable to load registration page", http.StatusInternalServerError)
			return
		}
		ParseAlertMessage(w, tmpl, "Invalid email too weak")
		return
	}

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
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	return regexp.MustCompile(regex).MatchString(email)

}
