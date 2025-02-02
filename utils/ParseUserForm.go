package utils

import (
	"forum/models"
	"net/http"
)

func PaerseUserForm(r *http.Request)(*models.User, *models.WebError){
	err := r.ParseForm()
	if err != nil{
		return nil, &models.WebError{Code: http.StatusBadRequest, Issue:"Invalid form submission"}
	}

	firstName := r.FormValue("first_name")
	lastName := r.FormValue("last_name")
	email := r.FormValue("email")
	username := r.FormValue("username")
	password1 := r.FormValue("password")
	AcceptedPassword := PasswordStrength(password1)
	bio := r.FormValue("bio")

	user := &models.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Username:  username,
		Password:  AcceptedPassword,
		Bio:  bio,
	}
	return user, nil
}
