package handlers

import (
	"net/http"
	"os"
)

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	// Restric requests to "GET"
	if r.Method != "GET" {
		BadRequestHandler(w)
		return
	}

	// Set predetermined path to catch malicious file tranversal
	filePath := "web/" + r.URL.Path
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
