package analyser

import (
	"fmt"

	"github.com/arctheowl/GoDeck/deck"
)

// CompareHands takes a board and two hands and returns the winner
func CompareHands(board []deck.Card, p1Cards []deck.Card, p2Cards []deck.Card) int {
	p1score, p1scoreString := Analyse(board, p1Cards)
	p2score, p2scoreString := Analyse(board, p2Cards)
	// If nothing else score = high card
	if p1score == p2score {
		fmt.Println(p1score, p2score)
		winner := CompareHighCard(p1Cards, p2Cards)
		switch winner {
		case 1:
			p1score = 1 + p1score
		case 2:
			p2score = 1 + p2score
		case 0:
			break
		}
	}

	fmt.Printf("Player 1: has a score of %d with %s\n", p1score, p1scoreString)
	fmt.Printf("Player 2: has a score of %d with %s\n", p2score, p2scoreString)

	if p1score > p2score {
		fmt.Println("Player 1 wins")
		return 1
	}
	if p1score < p2score {
		fmt.Println("Player 2 wins")
		return 2
	}
	fmt.Println("It's a draw")
	return 0
}

// Analyse takes a board and a hand and returns a score
func Analyse(board []deck.Card, pCards []deck.Card) (int, string) {
	//Analyse the hands dealt to see which one wins a game of poker
	pscore := 0
	FinalScoreString := ""

	//Check if the flop is a pair
	isPair, score, scoreString := PairToFullHouse(board, pCards)
	if isPair {
		pscore += score
		FinalScoreString = scoreString
		if scoreString == "Full House" || scoreString == "Four of a Kind" {
			return pscore, scoreString
		}
	}

	// Check if the flop is a straight
	straightRes, HighestStraightCard := IsStraight(board, pCards)
	if straightRes {
		pscore = 400 + HighestStraightCard
		FinalScoreString = "Straight"
	}

	//Check if the flop is a flush
	flushRes, HighestFlushCard, _ := IsFlush(board, pCards)
	if flushRes {
		pscore = 500 + HighestFlushCard
		FinalScoreString = "Flush"
	}

	//Check if the flop is a straight flush
	if straightRes && flushRes {
		pscore = 900 + HighestStraightCard
		FinalScoreString = "Straight Flush"
	}

	//Check if the flop is a royal flush
	if straightRes && flushRes && HighestStraightCard == 14 && HighestFlushCard == 14 {
		pscore = 1000
		FinalScoreString = "Royal Flush"
	}

	if pscore == 0 {
		pscore = IsHighCard(pCards)
		FinalScoreString = "High Card"
	}

	return pscore, FinalScoreString
}

/*Scoring
13*13 = 169

high card = value of highest unique card
pair = 100 + pairValue
	- Lowest value (Pair of 2)= 102
	- Highest value (Pair of Aces)= 114
Two pair = 150 + pairValue
	- Lowest value (Pair of 2)= 152
	- Highest value (Pair of Aces)= 164
Three of a Kind = 200 + KindValue
	- Lowest value (3 of 2)= 202
	- Highest value (3 of Ace)= 214
Straight = 400 + HighestStraightCard
	- Lowest value (Ace to 5)= 405
	- Highest value (10 to Ace)= 414
Flush = 500 + HighestFlushCard
	- Lowest value (Ace,2,3,4,6)= 506
	- Highest value (9, J, Q, K, Ace)= 514
Full House = 600 + 3KindValue
	- Lowest value (3 of 2, 2 of 3)= 506
	- Highest value (9, J, Q, K, Ace)= 514
Four of a Kind =  700 + 4KindValue
	- Lowest value (4 of 2)= 706
	- Highest value (4 of Ace)= 714
Straight Flush = 800 + HighestStraightCard
	- Lowest value (Ace to 5)= 805
	- Highest value (9 to K)= 813
Royal Flush = 1000
*/
