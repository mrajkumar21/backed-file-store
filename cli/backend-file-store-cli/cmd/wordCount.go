package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

type TotalWordCount struct {
	WordCount int
}

func init() {
	var frequentWord = &cobra.Command{
		Use:   "wc",
		Short: "Gives Total Number of words in the server",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("WordCount called")

			err := wordCount(args)
			if err == nil {
				fmt.Println("File Operation Success")
			} else {
				fmt.Println("File Operation Failed")
			}

		},
	}
	rootCmd.AddCommand(frequentWord)

}

func wordCount(files []string) error {
	client := &http.Client{}
	totalWordCount := TotalWordCount{}
	url := URL + "/wordcount"
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
	err = json.Unmarshal(res, &totalWordCount)
	if err != nil {
		fmt.Println("unmarshal Failed:", err)
		return err
	}
	fmt.Println("totalWordCount:", totalWordCount.WordCount)
	return nil
}
