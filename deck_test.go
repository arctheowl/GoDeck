package deck_test

import (
	"testing"

	deck "github.com/arctheowl/GoDeck"
)

// Test_DeckID Tests if the deck will produce an ID to be used
func Test_DeckID(t *testing.T) {
	d := deck.New()

	if d.ID.String() == "" {
		t.Error("Deck.New() didn't produce an ID for the deck")
	}

}


// Test_Shuffle tests if the deck is correctly shuffled after deck.shuffle() is invoked
func Test_Shuffle(t *testing.T) {
	d := deck.New()

	x, y, z := d.Cards[0], d.Cards[1], d.Cards[2]

	d.Shuffle()

	a, b, c := d.Cards[0], d.Cards[1], d.Cards[2]

	if a == x {
		if b == y {
			if c == z {
				t.Error("Deck Shuffle() didn't Shufle correctly")
			}
		}
	}

}

// These "Count" Tests are to ensure that when a method is called that the deck contains the correct amount of cards at the end
func Test_DeckCount(t *testing.T) {
	d := deck.New()

	if len(d.Cards) != 52 {
		t.Error("Deck.New() produced the incorrect number of Cards")
	}
}

// Test_POPDeckCount Ensures that when using .Pop() only pops a single card out of the deck
func Test_POPDeckCount(t *testing.T) {
	d := deck.New()
	d, x := d.Pop()
	_ = x

	if len(d.Cards) != 51 {
		t.Error("Deck.Pop() produced the incorrect number of Cards")
	}
}


// Test_POP2DeckCount Ensures that when using .Pop2() only pops a couple of cards out of the deck
func Test_POP2DeckCount(t *testing.T) {
	d := deck.New()
	d, x, y := d.Pop2()
	_, _ = x, y

	if len(d.Cards) != 50 {
		t.Error("Deck.Pop2() produced the incorrect number of Cards")
	}
}

// Test_POPNumDeckCount Ensures that when using .PopNum() only pops the correct number of cards out of the deck
func Test_POPNumDeckCount(t *testing.T) {
	d := deck.New()
	d, _ = d.PopNum(7)
	

	if len(d.Cards) != 45 {
		t.Error("Deck.Pop2() produced the incorrect number of Cards")
	}
}
