package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

type FileName struct {
	Files []string `json:"files"`
}

func init() {
	var deleteCmd = &cobra.Command{
		Use:   "rm",
		Short: "Delete command will remove files from server ",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("delete called")

			if len(args) == 0 {

				fmt.Println("got empty argument")

				return
			}
			err := DeleteFile(args)
			if err == nil {
				fmt.Println("File Operation Success")
			} else {
				fmt.Println("File Operation Failed")
			}

		},
	}
	rootCmd.AddCommand(deleteCmd)

}

func DeleteFile(files []string) error {
	client := &http.Client{}
	fileName := FileName{}

	fileName.Files = append(fileName.Files, files...)

	body, _ := json.Marshal(fileName)
	url := URL + "/delete"
	request, err := http.NewRequest("DELETE", url, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("delete request failed:", err)
		return err
	}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("delete client request failed:", err)
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Println("Delete Failed:", resp.Status)
		return err
	}
	res, _ := io.ReadAll(resp.Body)
	fmt.Println("Response status:", string(res))
	return nil
}
