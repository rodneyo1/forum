package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"forum/models"
)

var hitch models.WebError

// Serves Bad Request error page
func BadRequestHandler(w http.ResponseWriter) {
	// Construct absolute path to error.html
	tmplPath, err := GetTemplatePath("error.html")
	if err != nil {
		log.Printf("Could not find template file: %v", err)
		return
	}

	w.WriteHeader(http.StatusBadRequest) // Set page header

	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		log.Println("Template parsing failed:", err)
		http.Error(w, "Could not load template, error page unavailable", http.StatusInternalServerError)
		return
	}
	// Set parameters of error
	hitch.Code = http.StatusBadRequest
	hitch.Issue = "Bad Request!"

	// Set the Content-Type header for the response
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Execute bad request template, handle emerging errors
	err = tmpl.Execute(w, hitch)
	if err != nil {
		http.Error(w, "Could not execute error template, error page unavailable", http.StatusInternalServerError)
		log.Println("Error executing template: ", err)
	}
}

// Serves Internal Server Error page
func InternalServerErrorHandler(w http.ResponseWriter) {
	// Construct absolute path to error.html
	tmplPath, err := GetTemplatePath("error.html")
	if err != nil {
		log.Printf("Could not find template file: %v", err)
		return
	}

	w.WriteHeader(http.StatusInternalServerError) // Set page header

	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Could not load template, error page unavailable", http.StatusInternalServerError)
		return
	}

	// Set parameters of error
	hitch.Code = http.StatusInternalServerError
	hitch.Issue = "Internal Server Error!"

	// Set the Content-Type header for the response
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Execute internal server error template, handle emerging errors
	err = tmpl.Execute(w, hitch)
	if err != nil {
		http.Error(w, "Could not execute error template, error page unavailable", http.StatusInternalServerError)
		log.Println("Error executing template: ", err)
	}
}

// Serves Not Found error page
func NotFoundHandler(w http.ResponseWriter) {
	// Construct absolute path to error.html
	tmplPath, err := GetTemplatePath("error.html")
	if err != nil {
		http.Error(w, "Could not find template file", http.StatusInternalServerError)
		log.Println("Could not find template file: ", err)
		return
	}

	w.WriteHeader(http.StatusNotFound) // Set page header

	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Could not load template, error page unavailable", http.StatusInternalServerError)
		return
	}
	// Set parameters of error
	hitch.Code = http.StatusNotFound
	hitch.Issue = "Not Found!"

	// Set the Content-Type header for the response
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Execute not found error template, handle emerging errors
	err = tmpl.Execute(w, hitch)
	if err != nil {
		http.Error(w, "Could not execute error template, error page unavailable", http.StatusInternalServerError)
		log.Println("Error executing template: ", err)
	}
}

func GetTemplatePath(templateFile string) (string, error) {
	// catch empty template files
	if templateFile == "" {
		return "", fmt.Errorf("template file name cannot be empty")
	}

	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Traverse up until we find the project root
	dir := wd
	for {
		// Construct path to template, check if constructed path exists
		templatePath := filepath.Join(dir, "web", "templates", templateFile)
		if _, err := os.Stat(templatePath); err == nil {
			return templatePath, nil
		}

		// Move up one directory
		parent := filepath.Dir(dir)
		if parent == dir { // Stop if we reach the root
			break
		}
		dir = parent
	}

	return "", fmt.Errorf("template file not found: %s", templateFile)
}
