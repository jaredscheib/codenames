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
		g := server.NewGame()
		p := cli.RenderTurnPrompt(g.WhoseTurn)
		fmt.Println(p)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
