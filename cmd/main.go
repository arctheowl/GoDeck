package main

import (
	"fmt"

	deck "github.com/arctheowl/GoDeck"
)

//This is just a very simple function that prints out the GoDeck to show that it's working and can be imported.
func main(){
	d := deck.New()
	fmt.Println(d)
	d, c := d.PopNum(2)

	fmt.Println("NEW DECK", d)
	fmt.Println("THIS IS THE CARDS:",c)
}
