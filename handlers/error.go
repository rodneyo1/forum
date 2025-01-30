package handlers

import (
	"html/template"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("error.html")
	if err != nil {
		http.Error(w, "Failed to load Error template", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, r); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
