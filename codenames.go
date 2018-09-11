package codenames

type Codename struct {
	// Position   string `json:"position"`
	Type       string `json:"type"`
	Value      string `json:"value"`
	Visibility string `json:"visibility"`
}

type Player struct {
	ID string `json:"id"`
}

type Team struct {
	Color     string   `json:"color"`
	Spymaster Player   `json:"spymaster"`
	Teammates []Player `json:"teammates"`
}

type Teams struct {
	Red  Team `json:"red"`
	Blue Team `json:"blue"`
}

type Game struct {
	WhoseTurn string     `json:"whoseTurn"`
	Teams     Teams      `json:"teams"`
	Field     []Codename `json:"field"`
}

func main() {

}
