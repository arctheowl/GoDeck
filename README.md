# GoDeck

![Image of GopherDeck](https://github.com/arctheowl/GoDeck/blob/master/GoDeckLogo.png)


This Module is a deck of cards implementation using Golang. 


There are several functions to help you create a card game with using a Go backend.
This module contains two structs "Card" and "Deck".

The Deck has 2 fields, "Cards" which is a list of "Card" structs, and ID which is a uuid ID for those that would like multi deck games and to keep track of individual decks.
This UUID package is the only dependency in this module.

# Methods:

* New - Creates a New Deck
* Shuffle - Shuffles and existing deck
* Pop - returns the deck and then a card
* Pop2 - returns the deck and then 2 cards

# Tests:

I have also implemented several tests that you cna execute as soon as you download to ensure that the package is operating correctly from your first download.
To run these tests just navigate to the directory of the package and run the follow command:
```bash
go test tests/deck_test.go
```