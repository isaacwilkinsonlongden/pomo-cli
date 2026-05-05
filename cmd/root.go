// Package cmd provides command logic
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd is a cobra.Command that acts as the run command for the program
var rootCmd = &cobra.Command{
	Use:   "pomo",
	Short: "A Pomodoro timer for the terminal",
	Long:  "pomo is a CLI tool to help you stay focused using Pomodoro techniques",
}

// Execute starts the program by running the rootCmd and checks if an error is retured
// if an error is returned we print the error and exit with status code 1
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
