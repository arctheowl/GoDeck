package main

import (
	"fmt"

	"github.com/arctheowl/GoDeck/analyser"
	"github.com/arctheowl/GoDeck/deck"
)

//This is just a very simple function that prints out the GoDeck to show that it's working and can be imported.
func main() {
	d := deck.New()
	// fmt.Println(d)

	d.Shuffle()

	d, player1Cards := d.PopNum(2)
	d, player2Cards := d.PopNum(2)
	_, flop := d.PopNum(5)

	// fmt.Println("NEW DECK", d.Cards)
	fmt.Println("THIS IS P1 CARDS:", player1Cards)
	fmt.Println("THIS IS P2 CARDS:", player2Cards)
	fmt.Println("Flop:", flop)

	analyser.CompareHands(flop, player1Cards, player2Cards)

}
