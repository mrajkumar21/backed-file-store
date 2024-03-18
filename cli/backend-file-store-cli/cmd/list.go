package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

func init() {
	var listCmd = &cobra.Command{
		Use:   "ls",
		Short: "List command will show all the files in the server",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("list called")

			err := ListFile(args)
			if err == nil {
				fmt.Println("File Operation Success")
			} else {
				fmt.Println("File Operation Failed")
			}
		},
	}
	rootCmd.AddCommand(listCmd)

}

func ListFile(files []string) error {
	client := &http.Client{}
	fileName := FileName{}
	url := URL + "/list"
	response, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("List NewRequest error:", err)
		return err
	}
	resp, err := client.Do(response)
	if err != nil {
		fmt.Println("List Client creation error:", err)
		return err
	}
	defer resp.Body.Close()
	res, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		fmt.Println("List Failed:", resp.Status)
		return err
	}
	err = json.Unmarshal(res, &fileName)
	if err != nil {
		fmt.Println("unmarshal Failed:", err)
		return err
	}
	for _, fname := range fileName.Files {
		fmt.Println(fname)
	}
	return nil
}
