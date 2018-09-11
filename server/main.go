package server

import (
	"github.com/jaredscheib/codenames"
)

func NewGame() codenames.Game {
	// TODO: generate field of codenames
	game := codenames.Game{
		WhoseTurn: "Red",
		Teams: codenames.Teams{
			Red:  codenames.Team{Color: "Red"},
			Blue: codenames.Team{Color: "Blue"},
		},
		Field: []codenames.Codename{
			{
				Type:       "Assassin",
				Value:      "joe",
				Visibility: "hidden",
			},
			{
				Type:       "Bystander",
				Value:      "bob",
				Visibility: "visible",
			},
			{
				Type:       "Red",
				Value:      "helena",
				Visibility: "hidden",
			},
			{
				Type:       "Blue",
				Value:      "jimmy",
				Visibility: "visible",
			},
		},
	}
	return game
}
