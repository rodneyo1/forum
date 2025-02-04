package posts

import (
	"net/http"
	"html/template"
	"forum/database"
	"strconv"
	"log"
)

func DisplaySinglePost(w http.ResponseWriter, r *http.Request){
	tmpl,err:=template.ParseFiles("web/templates/post.html");
	if err!=nil{
		http.Error(w, "Failed to load post template", http.StatusInternalServerError)
		return
	}
	postID,err := strconv.Atoi(r.URL.Query().Get("id"))
	if err!=nil{
		http.Error(w, "Failed to get post ID", http.StatusInternalServerError)
		return
		}

	PostData,err:=database.GetPostByUUID(postID)
	
	if err!=nil{
		log.Println("Error getting post data: ", err)
		return
	}
	if err := tmpl.Execute(w, PostData); err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}