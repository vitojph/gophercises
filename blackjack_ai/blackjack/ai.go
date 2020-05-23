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

type dealerAI struct{}

func (ai dealerAI) Bet() int {
	return 1
}

func (ai dealerAI) Play(hand []deck.Card, dealer deck.Card) Move {
	dealerScore := Score(hand...)
	if dealerScore <= 16 || dealerScore == 17 && Soft(hand...) {
		return MoveHit
	}
	return MoveStand
}

func (ai dealerAI) Results(hand [][]deck.Card, dealer []deck.Card) {
}

func HumanIntelligence() AI {
	return humanIntelligence{}
}

type humanIntelligence struct{}

func (ai humanIntelligence) Bet() int {
	return 1
}

func (ai humanIntelligence) Play(hand []deck.Card, dealer deck.Card) Move {
	for {
		fmt.Println("Player:", hand)
		fmt.Println("Dealer:", dealer)
		fmt.Println("What will you do? (h)it, (s)tand?")
		var input string
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			return MoveHit
		case "s":
			return MoveStand
		default:
			fmt.Println("Invalid option, try again!")
		}
	}
}

func (ai humanIntelligence) Results(hand [][]deck.Card, dealer []deck.Card) {
	fmt.Println("== FINAL HANDS ==")
	fmt.Println("Player:", hand)
	fmt.Println("Dealer:", dealer)
}
