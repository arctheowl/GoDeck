# GoDeck

![Image of GopherDeck](https://github.com/arctheowl/GoDeck/blob/master/GoDeckLogo.png)

![Image of GopherDeck](https://github.com/arctheowl/GoDeck/blob/master/coverage_badge.png)

This Module is a deck of cards implementation using Golang. 

There are several functions to help you create a card game with using a Go backend.
This module contains two structs "Card" and "Deck".

The Deck has 2 fields, "Cards" which is a list of "Card" structs, and ID which is a uuid ID for those that would like multi deck games and to keep track of individual decks.
This UUID package is the only dependency in this module.

The Analyser has many different functions that allow you to determine the winner of a hand. The methods can also be used to determine the strength of a hand on it's own.
# Usage:

To import this module you have to first go get the module and then import the module into your project using the commands below.

```bash
go get -u "github.com/arctheowl/GoDeck"
```


```go
import (
    deck "github.com/arctheowl/GoDeck"
)
```



# Methods:

**Deck**
* New - Creates a New Deck
* Shuffle - Shuffles and existing deck
* Pop - returns the deck and then a card
* Pop2 - returns the deck and then 2 cards
* PopNum - returns the deck and then the passed number of cards

**Analyser** 
* CompareHands - Takes 3 deck.Card arrays, for board and two player's hands, returns Int of 1 if player1 won, 2 if player2 won and 0 if tied
* Analyse - Takes 2 deck.Card arrays, for board and a player's hands, returns player score and scoreString 
* IsStraight - Takes 2 deck.Card arrays, for board and a player's hands, returns boolean along with highest straight card if true.
* IsFlush - Takes 2 deck.Card arrays, for board and a player's hands, returns boolean along with highest flush card if true.
* PairToFullHouse - Takes 2 deck.Card arrays, for board and a player's hands, returns true if the cards are a pair, three of a kind, four of a kind, or fullhouse. Along with card.Value and hand strength string.
* CompareHighCard - Takes 2 deck.Card arrays, for 2 player's hands, compares two hands highCards and returns int indicating the winner
* IsHighCard - Takes 1 deck.Card arrays, for 1 player's hands, and returns an int for the highest card.Value in the hand.
* CompareSingleHighCard - Takes 2 single deck.Cards, returns 1 for first card higher, 2 for second card higher and 0 for draw.

# Tests:

I have also implemented several tests that you can execute as soon as you download to ensure that the package is operating correctly from your first download.
Currently this module has 100% test coverage.

To run these tests just navigate to the directory of the package and run the follow command:

```bash
go test
```
