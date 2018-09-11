package main

import (
	"fmt"

	"github.com/jaredscheib/codenames/cli"
	"github.com/jaredscheib/codenames/server"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use: "start",
	Run: func(cmd *cobra.Command, args []string) {
		game := server.NewGame()
		str := cli.RenderGame(game)
		fmt.Println(str)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
