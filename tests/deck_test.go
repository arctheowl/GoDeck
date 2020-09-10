package deck_test

import (
	"fmt"
	"testing"

	"github.com/arctheowl/GoDeck.git/deck"
)

func Test_DeckID(t *tseting.T){
	d := deck.New()

	if 

}


// These "Count" Tests are to ensure that when a method is called that the deck contains the correct amount of cards at the end
func Test_DeckCount(t *testing.T) {
	d := deck.New()

	if len(d.Cards) != 52 {
		t.Error("Deck.New() produced the incorrect number of Cards")
	}
}

func Test_POPDeckCount(t *testing.T) {
	d := deck.New()
	d, x := d.Pop()
	fmt.Println(x)

	if len(d.Cards) != 51 {
		t.Error("Deck.Pop() produced the incorrect number of Cards")
	}
}

func Test_POP2DeckCount(t *testing.T) {
	d := deck.New()
	d, x, y := d.Pop2()
	fmt.Println(x, y)

	if len(d.Cards) != 50 {
		t.Error("Deck.Pop2() produced the incorrect number of Cards")
	}
}
