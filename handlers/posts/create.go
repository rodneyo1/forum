package posts_test

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"forum/database"
	"forum/utils"
)

// Handler for serving the form and handling form submission
func PostCreate(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Serve the form for creating a post
		tmpl := template.Must(template.New("form").Parse(`
			<!DOCTYPE html>
			<html>
			<head>
				<title>Create Post</title>
			</head>
			<body>
				<h1>Create a New Post</h1>
				<form method="POST" action="/posts/create" enctype="multipart/form-data">
					<label for="title">Title:</label><br>
					<input type="text" id="title" name="title" required><br><br>
					<label for="content">Content:</label><br>
					<textarea id="content" name="content" required></textarea><br><br>
					<label for="media">Upload Image (PNG or JPG, max 20MB):</label><br>
					<input type="file" id="media" name="media" accept=".png,.jpg,.jpeg" required><br><br>
					<input type="submit" value="Submit">
				</form>
			</body>
			</html>
		`))
		tmpl.Execute(w, nil)

	case http.MethodPost:
		if err := r.ParseMultipartForm(20 << 20); err != nil { // 20MB max
			http.Error(w, "File upload too large or invalid form data", http.StatusBadRequest)
			return
		}

		// extract the form values
		title := r.FormValue("title")
		content := r.FormValue("content")

		// handle the uploaded file
		file, handler, err := r.FormFile("media")
		if err != nil {
			http.Error(w, "Failed to retrieve the file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		// validate the file extension type and size
		allowedTypes := map[string]bool{
			"image/png":  true,
			"image/jpeg": true,
		}
		fileType := handler.Header.Get("Content-Type")
		if !allowedTypes[fileType] {
			http.Error(w, "Invalid file type. Only PNG and JPG images are allowed.", http.StatusBadRequest)
			return
		}

		// Generate a random filename
		randomFileName, err := utils.GenerateRandomName()
		if err != nil {
			http.Error(w, "Failed to generate a unique filename", http.StatusInternalServerError)
			return
		}

		// determine the file extension based on the MIME type
		var ext string
		switch fileType {
		case "image/png":
			ext = ".png"
		case "image/jpeg":
			ext = ".jpg"
		}

		// construct the full filename
		filename := randomFileName + ext

		// save the file to the media folder
		// try create the media directory if it does not exist
		mediaFolder := "media"
		if err := os.MkdirAll(mediaFolder, os.ModePerm); err != nil {
			http.Error(w, "Failed to create media folder", http.StatusInternalServerError)
			return
		}

		filePath := filepath.Join(mediaFolder, filename)
		outFile, err := os.Create(filePath)
		if err != nil {
			http.Error(w, "Failed to save the file", http.StatusInternalServerError)
			return
		}
		defer outFile.Close()

		// copy the file content to the new file
		if _, err := io.Copy(outFile, file); err != nil {
			http.Error(w, "Failed to save the file", http.StatusInternalServerError)
			return
		}

		// mock user ID for now
		userID := 1

		tableName := "posts" // table name
		columns := []string{"user_id", "title", "content", "media"} // columns to insert
		values := []interface{}{userID, title, content, filename} // values to insert

		postID, err := database.Create(tableName, columns, values)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to create post: %v", err), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Post created successfully! Post ID: %d\n", postID)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
