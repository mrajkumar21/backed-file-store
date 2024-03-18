package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

var FrequentWordCheck struct {
	Items []struct {
		FrequencyWord string `json:"frequencyword"`
		WordCount     int    `json:"wordCount"`
	} `json:"items"`
}

func init() {
	var frequentWord = &cobra.Command{
		Use:   "freq-words",
		Short: "Get most frequently used words from the uploaded documentes",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("frequentWord called")

			limit, _ := cmd.Flags().GetString("limit")

			sort, _ := cmd.Flags().GetString("sort")

			err := frequentWordCheck(limit, sort)
			if err == nil {
				fmt.Println("File Operation Success")
			} else {
				fmt.Println("File Operation Failed")
			}

		},
	}
	rootCmd.AddCommand(frequentWord)
	frequentWord.Flags().StringP("limit", "l", "10", "limit the no of output")
	frequentWord.Flags().StringP("sort", "s", "d", "to get output sort in ascending or descending pass a for assending and d for decending")

}

func frequentWordCheck(limit string, sort string) error {
	client := &http.Client{}
	url := URL + "/frequentword/" + sort + "/" + limit
	response, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("frequentWordCheck Create Request error:", err)
		return err
	}

	resp, err := client.Do(response)
	if err != nil {
		fmt.Println("frequentWordCheck client creation error:", err)
		return err
	}
	if resp.StatusCode != 200 {
		fmt.Println("frequentWordCheck Failed:", err)
		return err
	}

	defer resp.Body.Close()

	res, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ReadAll Error:", err)
	}

	err = json.Unmarshal(res, &FrequentWordCheck)
	if err != nil {
		fmt.Println("unmarshal Failed:", err)
		return err
	}
	for _, item := range FrequentWordCheck.Items {
		fmt.Printf("%d   %s \n", item.FrequencyWord, item.WordCount)
	}
	return nil
}
