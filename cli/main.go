package cli

import (
	"fmt"

	"github.com/jaredscheib/codenames"
)

func RenderTurnPrompt(team string) string {
	str := fmt.Sprintf("%s Spymaster, it is your turn to offer a clue.", team)
	return str
}

func RenderField(g codenames.Game, view string) string {
	str := renderField(g.Field, view)
	return str
}

func renderField(f []codenames.Codename, view string) string {
	var str string
	for i, c := range f {
		if view == "spymaster" {
			str = str + fmt.Sprintf("%d. Type: %s, Value: %s, Visibility: %s", i+1, c.Type, c.Value, c.Visibility)
		} else {
			str = str + fmt.Sprintf("%d. Value: %s, Visibility: %s", i+1, c.Value, c.Visibility)
		}
		str = str + "\n"
	}
	return str
}
