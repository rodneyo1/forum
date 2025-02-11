package posts

import (
	errors "forum/handlers/errors"
	"html/template"
	"log"
	"net/http"

	"forum/database"
	"forum/models"
)

func PostDisplay(w http.ResponseWriter, r *http.Request) {
	loggedIn := false
	session, lIn := database.IsLoggedIn(r)
	if lIn {
		loggedIn = true
	}

	// Retrieve user data
	userData, _ := database.GetUserbySessionID(session.SessionID)

	tmpl, err := template.ParseFiles("./web/templates/post_display.html")
	if err != nil {
		log.Printf("ERROR: Could not parse template: %v", err)
		errors.InternalServerErrorHandler(w)
		return
	}
	postID := r.URL.Query().Get("pid")
	// fmt.Println("SINGLE PID: ", postID)

	postData, err := database.GetPostByUUID(postID)
	if err != nil {
		log.Println("Error getting post data: ", err)
		return
	}

	// Infuse data to be executed with inquiry if user is logged in
	data := struct {
		PostData models.PostWithCategories
		IsLogged bool
		ProfPic  string
	}{
		PostData: postData,
		IsLogged: loggedIn,
		ProfPic:  userData.Image,
	}

	// fmt.Println("POST: ", PostData)

	if err := tmpl.Execute(w, data); err != nil {
		log.Printf("Error executing template: %v", err)
		errors.InternalServerErrorHandler(w)
	}
}
