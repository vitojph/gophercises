package main

import (
	"fmt"

	"github.com/vitojph/gophercises/blackjack_ai/blackjack"
)

func main() {
	opts := blackjack.Options{Decks: 3, Hands: 2, BlackjackPayout: 1.5}
	game := blackjack.New(opts)
	winnings := game.Play(blackjack.HumanIntelligence())
	fmt.Println(winnings)
}
