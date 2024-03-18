package pkg

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

func CreateFile(files []*multipart.FileHeader) error {
	for _, fileHeader := range files {
		fmt.Println(UploadDirectory+"/"+fileHeader.Filename, " Creation Started")
		file, err := fileHeader.Open()
		if err != nil {
			fmt.Println("Unable to open the file:", err)
			return err
		}
		defer file.Close()

		// Create the destination file
		destination, err := os.Create(UploadDirectory + "/" + fileHeader.Filename)
		if err != nil {
			fmt.Println("Unable to create the file:", err)
			return err
		}
		defer destination.Close()

		// Copy the file to the destination
		if _, err := io.Copy(destination, file); err != nil {
			fmt.Println("Unable to copy the content:", err)
			return err
		}
		fmt.Println(UploadDirectory+"/"+fileHeader.Filename, " Creation Success")
	}
	return nil
}
