package main

import (
	"fmt"

	"github.com/jaredscheib/codenames/cli"
	"github.com/jaredscheib/codenames/server"
	"github.com/spf13/cobra"
)

var viewCmd = &cobra.Command{
	Use: "view",
	Run: func(cmd *cobra.Command, args []string) {
		g := server.NewGame() // TODO: make this the current game
		p := cli.RenderField(g, args[0])
		fmt.Println(p)
	},
}

func init() {
	rootCmd.AddCommand(viewCmd)
}
