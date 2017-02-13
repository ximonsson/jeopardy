package main

type Clue struct {
	Text string `json:"text"`
	Ans  string `json:"ans"`
}

type Category struct {
	Name  string `json:"name"`
	Clues []Clue `json:"clues"`
}

type Round []Category

type Game struct {
	Name    string
	Rounds  []Round
	Round   int
	FinalQ  Clue
	Players []Player
}

func Run() {
	for pid := range buzzers {
		if !buzzersActive {
			// ignore the player buzzers
			continue
		}
		currPlayer = players[pid]
		// wait for player to answer...
		<-answering
	}
}
