package main

import (
	"fmt"

	deck "github.com/arctheowl/GoDeck"
)

//This is just a very simple function that prints out the GoDeck to show that it's working and can be imported.
func main(){
	d := deck.New()
	fmt.Println(d)
	fmt.Println(d.PopNum(9))
	d = d.popNum(5)

}
