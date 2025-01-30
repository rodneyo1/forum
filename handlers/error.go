package handlers

import (
	"html/template"
	"log"
	"net/http"

	"forum/models"
)

var hitch models.WebError

// Serves Bad Request error page
func BadRequestHandler(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	tmpl, err := template.ParseFiles("web/templates/error.html")
	if err != nil {
		http.Error(w, "Could not load template, error page unavailable", http.StatusInternalServerError)
		return
	}
	// Set parameters of error
	hitch.Code = http.StatusBadRequest
	hitch.Issue = "Bad Request!"

	// Execute bad request template, handle emerging errors
	err = tmpl.Execute(w, hitch)
	if err != nil {
		http.Error(w, "Could not execute error template, error page unavailable", http.StatusInternalServerError)
		log.Println("Error executing template: ", err)
	}
}

// Serves Internal Server Error page
func InternalServerErrorHandler(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	tmpl, err := template.ParseFiles("web/templates/error.html")
	if err != nil {
		http.Error(w, "Could not load template, error page unavailable", http.StatusInternalServerError)
		return
	}

	// Set parameters of error
	hitch.Code = http.StatusInternalServerError
	hitch.Issue = "Internal Server Error!"

	// Execute internal server error template, handle emerging errors
	err = tmpl.Execute(w, hitch)
	if err != nil {
		http.Error(w, "Could not execute error template, error page unavailable", http.StatusInternalServerError)
		log.Println("Error executing template: ", err)
	}
}

// Serves Not Found error page
func NotFoundHandler(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	tmpl, err := template.ParseFiles("web/templates/error.html")
	if err != nil {
		http.Error(w, "Could not load template, error page unavailable", http.StatusInternalServerError)
		return
	}
	hitch.Code = http.StatusNotFound
	hitch.Issue = "Not Found!"

	// Execute not found error template, handle emerging errors
	err = tmpl.Execute(w, hitch)
	if err != nil {
		http.Error(w, "Could not execute error template, error page unavailable", http.StatusInternalServerError)
		log.Println("Error executing template: ", err)
	}
}
