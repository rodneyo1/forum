package handlers

import (
	"net/http"
	"os"
	"strings"
)

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	// Restric requests to "GET"
	if r.Method != "GET" {
		BadRequestHandler(w)
		return
	}

	path := r.URL.Path

	if strings.HasPrefix(path, "/static") {
		path = strings.TrimPrefix(path, "/static/")
	}

	// Set predetermined path to catch malicious file tranversal
	filePath := "./web/" + path
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		NotFoundHandler(w)
		return
	}

	// Restrict access to directories
	if fileInfo.IsDir() {
		NotFoundHandler(w)
		return
	}

	// Serve the requested file
	http.ServeFile(w, r, filePath)
}
