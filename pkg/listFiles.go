package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type FileName struct {
	Files []string `json:"files"`
}

func ListFiles(w http.ResponseWriter, r *http.Request) {
	dir, err := os.Open(UploadDirectory)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer dir.Close()

	// Read the directory contents
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Extract file names
	filesName := FileName{}
	for _, fileInfo := range fileInfos {
		if fileInfo.Mode().IsRegular() { // Check if it's a regular file
			filesName.Files = append(filesName.Files, fileInfo.Name())
			fmt.Println(fileInfo.Name())

		}
	}
	// Send the file list as a response
	w.Header().Set("Content-Type", "application/json")
	resBody, _ := json.Marshal(filesName)

	w.Write(resBody)
}
