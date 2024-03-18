package pkg

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
)

func Frequency(w http.ResponseWriter, r *http.Request) {
	Count := 0
	sortParam := r.URL.Query().Get("sort")
	limitParam := r.URL.Query().Get("limit")

	wordFrequency := make(map[string]int)

	// Open the directory
	dir, err := os.Open(UploadDirectory)
	if err != nil {
		fmt.Println("Error opening directory:", err)
		return
	}
	defer dir.Close()

	// Read the directory entries
	fileInfos, err := dir.Readdir(-1) // -1 means read all entries
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, fileInfo := range fileInfos {
		if fileInfo.Mode().IsRegular() { // Check if it's a regular file

			file, err := os.Open(UploadDirectory + "/" + fileInfo.Name())
			if err != nil {
				fmt.Println("Error opening file:", err)
				continue
			}
			defer file.Close()
			Scanner := bufio.NewScanner(file)
			Scanner.Split(bufio.ScanWords)
			for Scanner.Scan() {
				Count++
				wordFrequency[Scanner.Text()]++
			}
			if err := Scanner.Err(); err != nil {
				log.Fatal(err)
			}

		}
	}

	// Create a slice of WordFrequency objects
	var WordFrequencyList []struct {
		Word      string
		Frequency int
	}
	for word, frequency := range wordFrequency {
		WordFrequencyList = append(WordFrequencyList, struct {
			Word      string
			Frequency int
		}{word, frequency})
	}

	if sortParam == "a" {
		// Sort the slice in ascending order based on frequency
		sort.SliceStable(WordFrequencyList, func(i, j int) bool {
			return WordFrequencyList[i].Frequency < WordFrequencyList[j].Frequency
		})

	} else {
		// Sort the slice in descending order based on frequency
		sort.SliceStable(WordFrequencyList, func(i, j int) bool {
			return WordFrequencyList[i].Frequency > WordFrequencyList[j].Frequency
		})

	}

	nolimit, _ := strconv.Atoi(limitParam)

	if len(WordFrequencyList) < nolimit {

		nolimit = len(WordFrequencyList)

	}

	for i, wf := range WordFrequencyList[:nolimit] {
		fmt.Printf("%d. %s (%d occurrences)\n", i+1, wf.Word, wf.Frequency)
	}

	w.Header().Set("Content-Type", "application/json")
	resBody, _ := json.Marshal(WordFrequencyList)
	w.Write(resBody)
}
