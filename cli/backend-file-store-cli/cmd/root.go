package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	URL string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "store",
	Short: "Store is a cli tools to perform  file operations using command line. This cli can be used to upload file, deletefile, list the current files in the server.",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is cald by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
func init() {
	value := os.Getenv("STORE_URL")
	if value == "" {
		value = "http://localhost:8080"
	}
	URL = value

}
