package deck

import (
	"fmt"
	"math/rand"
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

func TestJoker(t *testing.T) {
	cards := New(Jokers(3))
	count := 0
	for _, card := range cards {
		if card.Suit == Joker {
			count++
		}
	}
	if count != 3 {
		t.Error("Expected 3 Jokers, but we found", count)
	}
}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}
	cards := New(Filter(filter))
	for _, c := range cards {
		if c.Rank == Two || c.Rank == Three {
			t.Error("Expected all 2s and 3s to be filtered out.")
		}
	}
}

func TestDeck(t *testing.T) {
	cards := New(Deck(3))
	// 3 decks * 13 ranks * 4 suits
	if len(cards) != 3*13*4 {
		t.Errorf("Expected %d cards, but received %d", 3*13*4, len(cards))
	}
}

func TestShuffle(t *testing.T) {
	// make shuffleRand deterministic
	shuffleRand = rand.New(rand.NewSource(0))
	orig := New()
	// we know these are the two first cards
	first := orig[40]
	second := orig[35]
	cards := New(Shuffle)
	if cards[0] != first {
		t.Errorf("Expected the first card to be %s, received %s.", first, cards[0])
	}
	if cards[1] != second {
		t.Errorf("Expected the second card to be %s, received %s.", second, cards[1])
	}
}
