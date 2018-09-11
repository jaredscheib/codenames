package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version number of Codenames",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Codenames v0.0.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
