package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	addFileCmd := &cobra.Command{
		Use:   "add",
		Short: "add command upload files to server",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("addFile called")

			if len(args) == 0 {

				fmt.Println("got empty argument")

				return
			}

			err := File(args)
			if err == nil {
				fmt.Println("File Operation Success")
			} else {
				fmt.Println("File Operation Failed")
			}

		},
	}
	rootCmd.AddCommand(addFileCmd)

}

func File(files []string) error {
	client := &http.Client{}
	for _, filePath := range files {
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return err
		}
		content, err := io.ReadAll(file)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return err
		}

		files := []struct {
			FileName    string
			FileContent string
		}{
			{FileName: file.Name(), FileContent: string(content)},
		}

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		for _, file := range files {
			part, _ := writer.CreateFormFile("files", file.FileName)
			part.Write([]byte(file.FileContent))
		}
		writer.Close()

		url := URL + "/add"

		request, err := http.NewRequest("POST", url, body)
		request.Header.Set("Content-Type", writer.FormDataContentType())

		if err != nil {
			fmt.Println("create request failed:", err)
			return err
		}
		response, err := client.Do(request)
		if err != nil {
			fmt.Println("client request failed:", err)
			return err
		}
		defer response.Body.Close()
		if response.StatusCode != 200 {
			fmt.Println("Create Failed:", response.Status)
			return errors.New("Invalid StatusCode received")
		}
		res, _ := io.ReadAll(response.Body)
		fmt.Println("Response status:", string(res))
	}
	return nil

}
