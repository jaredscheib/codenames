package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "codenames",
	Short: "Codenames is a game",
	Run: func(cmd *cobra.Command, arg []string) {
		fmt.Println("Welcome to Codenames")
	},
}

// Execute runs the root cobra command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
