package handlers

import (
	"html/template"
	"log"
	"net/http"

	"forum/database"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Construct absolute path to login.html
	templatePath, err := GetTemplatePath("login.html")
	if err != nil {
		InternalServerErrorHandler(w)
		log.Println("Could not find template file: ", err)
		return
	}

	// When method is GET, render form
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles(templatePath)
		if err != nil {
			http.Error(w, "Failed to load Login template", http.StatusInternalServerError)
			return
		}

		if err := tmpl.Execute(w, r); err != nil {
			InternalServerErrorHandler(w)
			return
		}
		return
	}

	// Restrict requests to "GET" or "POST"
	if r.Method != "POST" {
		BadRequestHandler(w)
		log.Println("LoginHandler ERROR: Bad request")
		return
	}

	// Parse form data from the request
	err = r.ParseForm()
	if err != nil {
		BadRequestHandler(w)
		log.Println("Invalid form submission", http.StatusBadRequest)
		return
	}

	// Extract form data
	emailUsername := r.FormValue("email_username")
	password := r.FormValue("password")

	// Validate credentials
	if !database.VerifyUser(emailUsername, password) {
		ParseAlertMessage(w, templatePath, "Invalid Username or Password") // Parse error message
		return
	}

	// Redirect user to home page
	http.Redirect(w, r, "/success", http.StatusFound)
}

func SuccessHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		BadRequestHandler(w)
		return
	}

	http.Error(w, "Loged in1!", http.StatusInternalServerError)
}

func ParseAlertMessage(w http.ResponseWriter, tmplPath, message string) {
	// Define template path and error message
	alert := map[string]string{"ErrorMessage": message}

	// Render page
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		InternalServerErrorHandler(w)
		log.Printf("Could not parse template: %v", err)
		return
	}

	// Execute page
	err = tmpl.Execute(w, alert)
	if err != nil {
		InternalServerErrorHandler(w)
		log.Printf("Could not execute template %v", err)
		return
	}
}
