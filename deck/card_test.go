package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Two, Suit: Spade})
	fmt.Println(Card{Rank: Nine, Suit: Diamond})
	fmt.Println(Card{Rank: Jack, Suit: Club})
	fmt.Println(Card{Suit: Joker})

	// Output:
	// Ace of Hearts
	// Two of Spades
	// Nine of Diamonds
	// Jack of Clubs
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()
	if len(cards) != 13*4 {
		t.Error("Wrong number of cards in the new deck.")
	}

}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	expected1stcard := Card{Rank: Ace, Suit: Spade}
	if cards[0] != expected1stcard {
		t.Error("Expected Ace of Spades as first card. Received:", cards[0])
	}
}

func TestSort(t *testing.T) {
	cards := New(Sort(Less))
	expected1stcard := Card{Rank: Ace, Suit: Spade}
	if cards[0] != expected1stcard {
		t.Error("Expected Ace of Spades as first card. Received:", cards[0])
	}
}
