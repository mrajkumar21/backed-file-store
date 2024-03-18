package pkg

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type TotalWordCount struct {
	WordCount int
}

func WordCount(w http.ResponseWriter, r *http.Request) {

	count := 0
	totaWordCount := TotalWordCount{}
	// Open the directory
	file, err := os.Open(UploadDirectory)
	if err != nil {
		fmt.Println("Error opening directory:", err)
		return
	}
	defer file.Close()

	// Read the directory entries
	fileInfos, err := file.Readdir(-1) // -1 means read all entries
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
				count++
			}
			if err := Scanner.Err(); err != nil {
				log.Fatal(err)
			}
		}
	}
	totaWordCount.WordCount = count
	w.Header().Set("Content-Type", "application/json")
	resBody, _ := json.Marshal(totaWordCount)
	w.Write(resBody)
}
