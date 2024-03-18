package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type FileBody struct {
	Files []string `json:"files"`
}

func DeleteFile(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	FileBodyreq := FileBody{}

	json.Unmarshal(body, &FileBodyreq)

	// Remove the file
	for _, fileName := range FileBodyreq.Files {
		filePath := UploadDirectory + "/" + fileName
		err := os.Remove(filePath)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to delete file: %s", err.Error()), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "File %s deleted successfully", fileName)

	}

	// Respond to the client indicating success
	w.WriteHeader(http.StatusOK)
}
