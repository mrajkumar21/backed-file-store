package pkg

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func UpdateFile(w http.ResponseWriter, r *http.Request) {
	files := r.MultipartForm.File["files"]
	for _, fileHeader := range files {
		fmt.Println("UpdateFile Started for:", UploadDirectory+fileHeader.Filename)
		file, err := fileHeader.Open()
		if err != nil {
			fmt.Println("Unable to open the file:", err)
			return
		}
		defer file.Close()

		if _, err := os.Stat(UploadDirectory + "/" + fileHeader.Filename); os.IsExist(err) {
			err := os.Remove(UploadDirectory + "/" + fileHeader.Filename)
			if err != nil {
				fmt.Println("Unable to remove the file:", err)
				http.Error(w, fmt.Sprintf("Failed to delete file: %s", err.Error()), http.StatusInternalServerError)
				return
			}
		}
		// Save the file to disk
		f, err := os.Create(UploadDirectory + "/" + fileHeader.Filename)
		if err != nil {
			fmt.Println("Create Error:", err)
			return
		}
		defer f.Close()

		// Copy file data to the created file
		_, err = io.Copy(f, file)
		if err != nil {
			fmt.Println("Copy Error:", err)
			return
		}

		w.Write([]byte("File uploaded successfully"))
	}
}
