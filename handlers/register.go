package handlers

import (
	"html/template"
	"net/http"
)

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the method is allowed (e.g., POST)
	if !MethodCheck(w, r, http.MethodPost) {
		// If the method is not allowed, stop further processing
		return
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
