package main

import (
	"fmt"
	"net/http"
	"store-project/pkg"
)

func main() {
	http.HandleFunc("/add", pkg.AddFiles)
	http.HandleFunc("/list", pkg.ListFiles)
	http.HandleFunc("/delete", pkg.DeleteFile)
	http.HandleFunc("/update", pkg.UpdateFile)
	http.HandleFunc("/wordcount", pkg.WordCount)
	http.HandleFunc("/frequentword/:sort/:limit", pkg.Frequency)

	port := ":8080"
	fmt.Printf("Server listening on port %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Printf("Failed to start server: %s\n", err)
	}
}
