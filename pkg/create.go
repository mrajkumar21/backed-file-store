package pkg

import (
	"fmt"
	"net/http"
	"os"
)

func AddFiles(w http.ResponseWriter, r *http.Request) {
	// Parse the multipart form data
	err := r.ParseMultipartForm(10 << 20) // Limit to 10 MB
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Retrieve the files from the request
	files := r.MultipartForm.File["files"]

	// Create the upload directory if it doesn't exist
	if _, err := os.Stat(UploadDirectory); os.IsNotExist(err) {
		os.Mkdir(UploadDirectory, 0755)
	}

	// Iterate through each file and save it to the server
	err = CreateFile(files)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Files uploaded successfully")
}
