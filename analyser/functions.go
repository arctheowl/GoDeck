package analyser

import (
	"sort"

	"github.com/arctheowl/GoDeck/deck"
)

// IsStraight returns true if the cards are a straight, and the highestCard value
func IsStraight(board []deck.Card, playerCards []deck.Card) (bool, int) {
	//Check if the cards are a straight
	occured := map[int]bool{}
	result := []int{}

	for e := range playerCards {
		if !occured[playerCards[e].Value] {
			result = append(result, playerCards[e].Value)
			occured[playerCards[e].Value] = true
		}

	}

	for e := range board {
		if !occured[board[e].Value] {
			result = append(result, board[e].Value)
			occured[board[e].Value] = true
		}
	}

	for i := range result {

		if occured[result[i]+1] && occured[result[i]+2] && occured[result[i]+3] && occured[result[i]+4] {
			return true, result[i] + 4
		}

	}

	if occured[1] && occured[10] && occured[11] && occured[12] && occured[13] {
		return true, 14
	}

	return false, 0
}

// IsFlush returns true if the cards are a flush
func IsFlush(board []deck.Card, playerCards []deck.Card) (bool, int, string) {

	var highestClub int
	var highestDiamond int
	var highestHeart int
	var highestSpade int

	var clubCount int
	var diamondCount int
	var heartCount int
	var spadeCount int

	for _, card := range board {
		switch card.Suit {
		case "c":
			clubCount++
		case "d":
			diamondCount++
		case "h":
			heartCount++
		case "s":
			spadeCount++
		}
	}

	if clubCount < 3 && diamondCount < 3 && heartCount < 3 && spadeCount < 3 {
		return false, 0, ""
	}

	//Check if the cards are a flush
	for _, card := range playerCards {
		if card.Value == 1 {
			card.Value = 14
		}
		switch card.Suit {
		case "c":
			if card.Value > highestClub {
				highestClub = card.Value
			}
			clubCount++
		case "d":
			if card.Value > highestDiamond {
				highestDiamond = card.Value
			}
			diamondCount++
		case "h":
			if card.Value > highestHeart {
				highestHeart = card.Value
			}
			heartCount++
		case "s":
			if card.Value > highestSpade {
				highestSpade = card.Value
			}
			spadeCount++
		}
	}

	if clubCount >= 5 {
		return true, highestClub, "c"
	} else if diamondCount >= 5 {
		return true, highestDiamond, "d"
	} else if heartCount >= 5 {
		return true, highestHeart, "h"
	} else if spadeCount >= 5 {
		return true, highestSpade, "s"
	}
	return false, 0, ""
}

// PairToFullHouse returns true if the cards are a pair, but also checks to see if they are three of a kind, four of a kind, or fullhouse
func PairToFullHouse(board []deck.Card, playerCards []deck.Card) (bool, int, string) {
	var values = make(map[int]int)
	strength := 0
	strengthString := "No Pair"

	for _, card := range playerCards {
		values[card.Value]++
	}

	for _, card := range board {
		values[card.Value]++
	}

	for cardValue, value := range values {
		if value == 2 {
			strength = 100 + cardValue
			delete(values, cardValue)
			strengthString = "Pair"
			break
		} else if value == 3 {
			strength = 200 + cardValue
			delete(values, cardValue)
			strengthString = "Three of a Kind"
			break
		} else if value == 4 {
			strength = 700 + cardValue
			strengthString = "Four of a Kind"
		}
	}

	if strengthString == "Pair" {
		for cardValue, value := range values {
			if value == 2 {
				strength += 50 + cardValue
				strengthString = "Two Pair"
			}
			if value == 3 {
				strength = 600 + cardValue
				strengthString = "Full House"
			}
		}
	} else if strengthString == "Three of a Kind" {
		for cardValue, value := range values {
			if value == 2 {
				strength = 600 + cardValue
				strengthString = "Full House"
			}
		}
	}

	if strength == 0 {
		return false, 0, strengthString
	}

	return true, strength, strengthString
}

// CompareHand compares two hands highCards and returns the winner
func CompareHighCard(p1Cards []deck.Card, p2Cards []deck.Card) int {
	p1HighCard := IsHighCard(p1Cards)
	p2HighCard := IsHighCard(p2Cards)

	if p1HighCard > p2HighCard {
		return 1
	} else if p1HighCard < p2HighCard {
		return 2
	}

	finalReturn := CompareSingleHighCard(p1Cards[0], p2Cards[0])

	return finalReturn
}

// IsHighCard CardValue of the unique high card
func IsHighCard(playerCards []deck.Card) int {
	var values = []int{}
	for _, card := range playerCards {
		values = append(values, card.Value)
	}
	sort.Ints(values)

	if values[1] == 1 || values[0] == 1 {
		return 14
	}

	return values[1]
}

// CompareSingleHighCard compares to sinlge cards and returns int indicating higher card
func CompareSingleHighCard(p1Card deck.Card, p2Card deck.Card) int {
	new1Value := p1Card.Value
	new2Value := p2Card.Value

	if p1Card.Value == 1 {
		new1Value = 14
	}
	if p2Card.Value == 1 {
		new2Value = 14
	}

	if new1Value > new2Value {
		return 1
	}
	if new1Value < new2Value {
		return 2
	}
	return 0
}
