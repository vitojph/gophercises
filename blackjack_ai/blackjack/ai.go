package blackjack

import (
	"fmt"

	"github.com/vitojph/gophercises/deck"
)

type AI interface {
	Bet(shuffled bool) int
	Play(hand []deck.Card, dealer deck.Card) Move
	Results(hand [][]deck.Card, dealer []deck.Card)
}

type dealerAI struct{}

func (ai dealerAI) Bet(shuffled bool) int {
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

func (ai humanIntelligence) Bet(shuffled bool) int {
	if shuffled {
		fmt.Println("The deck was just shuffled")
	}
	fmt.Println("What would you like to bet?")
	var bet int
	fmt.Scanf("%d\n", &bet)
	return bet
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
