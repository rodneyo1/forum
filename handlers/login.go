package handlers

import (
	"html/template"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/login.html")
	if err != nil {
		http.Error(w, "Failed to load Login template", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, r); err != nil {
		InternalServerErrorHandler(w)
		return
	}
}
