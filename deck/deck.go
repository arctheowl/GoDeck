package deck

import (
	"math/rand"
	"time"
)

// Card struct is a represenation of a playing card, a value will be an int and the suit is a single letter string to represent the 4 suited deck
type Card struct {
	Value int
	Suit  string
}

//Deck is a list of Cards using the "Card" struct which can be found in the card.go file
type Deck struct {
	Cards []Card
}

//New - This will return a New shuffled deck
func New() Deck {

	d := Deck{}
	for suit := 0; suit < 4; suit++ {

		switch suit {
		case 0:
			for value := 1; value < 14; value++ {
				card1 := Card{value, "h"}
				d.Cards = append(d.Cards, card1)
			}
		case 1:
			for value := 1; value < 14; value++ {
				card1 := Card{value, "d"}
				d.Cards = append(d.Cards, card1)
			}
		case 2:
			for value := 1; value < 14; value++ {
				card1 := Card{value, "s"}
				d.Cards = append(d.Cards, card1)
			}
		case 3:
			for value := 1; value < 14; value++ {
				card1 := Card{value, "c"}
				d.Cards = append(d.Cards, card1)
			}
		}

	}
	d = d.Shuffle()
	return d
}

//Shuffle is a method that will shuffle any deck given using the random unix of the machine it's running on
func (d Deck) Shuffle() Deck {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.Cards), func(i, j int) { d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i] })

	return d
}

//Pop is a function that will return a card and the deck with that card removed so that you are able to deal a card to players
func (d Deck) Pop() (Deck, Card) {
	c := d.Cards[0]
	d.Cards = d.Cards[1:]

	return d, c
}

//Pop2 is a function that will return a card and the deck with that card removed so that you are able to deal a card to players
func (d Deck) Pop2() (Deck, Card, Card) {
	x, y := d.Cards[0], d.Cards[1]
	d.Cards = d.Cards[2:]

	return d, x, y
}
