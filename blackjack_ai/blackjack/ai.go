package blackjack

import (
	"fmt"

	"github.com/vitojph/gophercises/deck"
)

type AI interface {
	Bet() int
	Play(hand []deck.Card, dealer deck.Card) Move
	Results(hand [][]deck.Card, dealer []deck.Card)
}

type HumanIntelligence struct{}

func (ai *HumanIntelligence) Bet() int {
	return 1
}

func (ai *HumanIntelligence) Play(hand []deck.Card, dealer deck.Card) Move {
	for {
		fmt.Println("Player:", hand)
		fmt.Println("Dealer:", dealer)
		fmt.Println("What will you do? (h)it, (s)tand?")
		var input string
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			Hit
		case "s":
			Stand
		default:
			fmt.Println("Invalid option, try again!")
		}
	}
}

func (ai *HumanIntelligence) Results(hand [][]deck.Card, dealer []deck.Card) {}
	fmt.Println("== FINAL HANDS ==")
	fmt.Println("Player:", hand)
	fmt.Println("Dealer:", dealer)
}


type Move func(GameState) GameState

type GameState struct{}

func Hit(gs GameState) GameState {
	return gs
}

func Stand(gs GameState) GameState {
	return gs
}
