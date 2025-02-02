package handlers

import (
	"html/template"
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

	if err := tmpl.Execute(w, r); err != nil {
		InternalServerErrorHandler(w)
		return
	}
}
