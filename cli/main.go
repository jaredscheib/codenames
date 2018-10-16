package cli

import (
	"fmt"

	"github.com/jaredscheib/codenames"
)

func RenderGame(g codenames.Game) string {
	var str string
	str = str + fmt.Sprintf("%s Spymaster, it is your turn to offer a clue.", g.WhoseTurn)
	str = str + "\n"
	str = str + renderField(g.Field)
	return str
}

func renderField(f []codenames.Codename) string {
	var str string
	for _, c := range f {
		str = str + fmt.Sprintf("Type: %s, Value: %s, Visibility: %s", c.Type, c.Value, c.Visibility)
		str = str + "\n"
	}
	return str
}
