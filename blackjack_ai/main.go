package main

import (
	"fmt"
	
	"github.com/vitojph/gophercises/blackjack_ai/blackjack"
)

func main() {
	game := blackjack.New()
	winnings := game.Play(blackjack.HumanIntelligence())
	fmt.Println(winnings)
}
