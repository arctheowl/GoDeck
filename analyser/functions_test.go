package analyser

import (
	"testing"

	"github.com/arctheowl/GoDeck/deck"
)

// TestIsStraight tests the IsHigh function
func TestIsHighCard(t *testing.T) {

	cases := []struct {
		pCards         []deck.Card
		expectedResult int
	}{
		{
			pCards:         []deck.Card{{Name: "h1", Value: 1, Suit: "h"}, {Name: "h2", Value: 2, Suit: "h"}},
			expectedResult: 14,
		},
		{
			pCards:         []deck.Card{{Name: "c13", Value: 13, Suit: "c"}, {Name: "h1", Value: 1, Suit: "h"}},
			expectedResult: 14,
		},
		{
			pCards:         []deck.Card{{Name: "h10", Value: 10, Suit: "h"}, {Name: "h6", Value: 6, Suit: "h"}},
			expectedResult: 10,
		},
	}

	for _, c := range cases {
		result := IsHighCard(c.pCards)
		if result != c.expectedResult {
			t.Errorf("Expected %d, got %d", c.expectedResult, result)
		}
	}
}

// TestIsStraight tests the CompareHighCard function
func TestCompareHighCard(t *testing.T) {
	cases := []struct {
		p1Cards        []deck.Card
		p2Cards        []deck.Card
		expectedResult int
	}{
		{
			p1Cards:        []deck.Card{{Name: "h1", Value: 1, Suit: "h"}, {Name: "h2", Value: 2, Suit: "h"}},
			p2Cards:        []deck.Card{{Name: "h1", Value: 1, Suit: "h"}, {Name: "h2", Value: 2, Suit: "h"}},
			expectedResult: 0,
		},
		{
			p1Cards:        []deck.Card{{Name: "c13", Value: 13, Suit: "c"}, {Name: "h3", Value: 3, Suit: "h"}},
			p2Cards:        []deck.Card{{Name: "c12", Value: 12, Suit: "c"}, {Name: "h1", Value: 1, Suit: "h"}},
			expectedResult: 2,
		},
		{
			p1Cards:        []deck.Card{{Name: "h3", Value: 3, Suit: "h"}, {Name: "h4", Value: 4, Suit: "h"}},
			p2Cards:        []deck.Card{{Name: "h2", Value: 2, Suit: "h"}, {Name: "h4", Value: 4, Suit: "h"}},
			expectedResult: 1,
		},
		{
			p1Cards:        []deck.Card{{Name: "h13", Value: 13, Suit: "h"}, {Name: "h4", Value: 4, Suit: "h"}},
			p2Cards:        []deck.Card{{Name: "h2", Value: 2, Suit: "h"}, {Name: "h4", Value: 4, Suit: "h"}},
			expectedResult: 1,
		},
	}

	for _, c := range cases {
		result := CompareHighCard(c.p1Cards, c.p2Cards)
		if result != c.expectedResult {
			t.Errorf("Expected %d, got %d", c.expectedResult, result)
		}
	}

}

// TestCompareSingleHighCard tests the CompareSingleHighCard function
func TestCompareSingleHighCard(t *testing.T) {
	cases := []struct {
		p1Cards        deck.Card
		p2Cards        deck.Card
		expectedResult int
	}{
		{
			p1Cards:        deck.Card{Name: "h1", Value: 1, Suit: "h"},
			p2Cards:        deck.Card{Name: "h1", Value: 1, Suit: "h"},
			expectedResult: 0,
		},
		{
			p1Cards:        deck.Card{Name: "h13", Value: 13, Suit: "h"},
			p2Cards:        deck.Card{Name: "h1", Value: 1, Suit: "h"},
			expectedResult: 2,
		},
		{
			p1Cards:        deck.Card{Name: "h1", Value: 1, Suit: "h"},
			p2Cards:        deck.Card{Name: "h12", Value: 12, Suit: "h"},
			expectedResult: 1,
		},
	}

	for _, c := range cases {
		result := CompareSingleHighCard(c.p1Cards, c.p2Cards)
		if result != c.expectedResult {
			t.Errorf("Expected %d, got %d", c.expectedResult, result)
		}
	}

}

func TestIsFlush(t *testing.T) {
	cases := []struct {
		board          []deck.Card
		pCards         []deck.Card
		expectedResult ResultOfIsFlush
	}{
		{
			board:          []deck.Card{{Name: "h5", Value: 5, Suit: "h"}, {Name: "h6", Value: 6, Suit: "h"}, {Name: "h7", Value: 7, Suit: "h"}, {Name: "h13", Value: 13, Suit: "h"}, {Name: "h10", Value: 10, Suit: "h"}},
			pCards:         []deck.Card{{Name: "h3", Value: 3, Suit: "h"}, {Name: "h2", Value: 2, Suit: "h"}},
			expectedResult: ResultOfIsFlush{isFlush: true, HighestFlushCard: 3},
		},
		{
			board:          []deck.Card{{Name: "c5", Value: 5, Suit: "c"}, {Name: "s6", Value: 6, Suit: "s"}, {Name: "s7", Value: 7, Suit: "s"}, {Name: "c13", Value: 13, Suit: "c"}, {Name: "s10", Value: 10, Suit: "s"}},
			pCards:         []deck.Card{{Name: "h3", Value: 3, Suit: "h"}, {Name: "h2", Value: 2, Suit: "h"}},
			expectedResult: ResultOfIsFlush{isFlush: false, HighestFlushCard: 0},
		},
		{
			board:          []deck.Card{{Name: "d5", Value: 5, Suit: "d"}, {Name: "d6", Value: 6, Suit: "d"}, {Name: "d7", Value: 7, Suit: "d"}, {Name: "c13", Value: 13, Suit: "c"}, {Name: "s10", Value: 10, Suit: "s"}},
			pCards:         []deck.Card{{Name: "d8", Value: 8, Suit: "d"}, {Name: "d2", Value: 2, Suit: "d"}},
			expectedResult: ResultOfIsFlush{isFlush: true, HighestFlushCard: 8},
		},
		{
			board:          []deck.Card{{Name: "c5", Value: 5, Suit: "c"}, {Name: "c6", Value: 6, Suit: "c"}, {Name: "s7", Value: 7, Suit: "s"}, {Name: "c12", Value: 12, Suit: "c"}, {Name: "s10", Value: 10, Suit: "s"}},
			pCards:         []deck.Card{{Name: "c3", Value: 3, Suit: "c"}, {Name: "c2", Value: 2, Suit: "c"}},
			expectedResult: ResultOfIsFlush{isFlush: true, HighestFlushCard: 3},
		},
		{
			board:          []deck.Card{{Name: "s5", Value: 5, Suit: "s"}, {Name: "s6", Value: 6, Suit: "s"}, {Name: "s7", Value: 7, Suit: "s"}, {Name: "c13", Value: 13, Suit: "c"}, {Name: "s10", Value: 10, Suit: "s"}},
			pCards:         []deck.Card{{Name: "s10", Value: 10, Suit: "s"}, {Name: "h2", Value: 2, Suit: "h"}},
			expectedResult: ResultOfIsFlush{isFlush: true, HighestFlushCard: 10},
		},
	}

	for _, c := range cases {
		result, highestCard, _ := IsFlush(c.board, c.pCards)
		fullResult := ResultOfIsFlush{isFlush: result, HighestFlushCard: highestCard}

		if fullResult != c.expectedResult {
			t.Errorf("Expected %v, got %v", c.expectedResult.isFlush, fullResult.isFlush)
			t.Errorf("Expected %v, got %v", c.expectedResult.HighestFlushCard, fullResult.HighestFlushCard)
		}

	}
}

func TestIsStraight(t *testing.T) {
	cases := []struct {
		board          []deck.Card
		pCards         []deck.Card
		expectedResult ResultOfIsStraight
	}{
		{
			board:          []deck.Card{{Name: "h5", Value: 5, Suit: "h"}, {Name: "h6", Value: 6, Suit: "h"}, {Name: "h7", Value: 7, Suit: "h"}, {Name: "h13", Value: 13, Suit: "h"}, {Name: "h10", Value: 10, Suit: "h"}},
			pCards:         []deck.Card{{Name: "h3", Value: 3, Suit: "h"}, {Name: "h2", Value: 2, Suit: "h"}},
			expectedResult: ResultOfIsStraight{IsStraight: false, HighestStraightCard: 0},
		},
		{
			board:          []deck.Card{{Name: "h5", Value: 5, Suit: "h"}, {Name: "h6", Value: 6, Suit: "h"}, {Name: "h7", Value: 7, Suit: "h"}, {Name: "h13", Value: 13, Suit: "h"}, {Name: "h10", Value: 10, Suit: "h"}},
			pCards:         []deck.Card{{Name: "h3", Value: 3, Suit: "h"}, {Name: "h4", Value: 4, Suit: "h"}},
			expectedResult: ResultOfIsStraight{IsStraight: true, HighestStraightCard: 7},
		},
		{
			board:          []deck.Card{{Name: "h5", Value: 5, Suit: "h"}, {Name: "c4", Value: 4, Suit: "c"}, {Name: "s3", Value: 3, Suit: "s"}, {Name: "h13", Value: 13, Suit: "h"}, {Name: "h10", Value: 10, Suit: "h"}},
			pCards:         []deck.Card{{Name: "h1", Value: 1, Suit: "h"}, {Name: "h2", Value: 2, Suit: "h"}},
			expectedResult: ResultOfIsStraight{IsStraight: true, HighestStraightCard: 5},
		},
		{
			board:          []deck.Card{{Name: "h11", Value: 11, Suit: "h"}, {Name: "c13", Value: 13, Suit: "c"}, {Name: "s3", Value: 3, Suit: "s"}, {Name: "h13", Value: 13, Suit: "h"}, {Name: "s12", Value: 12, Suit: "s"}},
			pCards:         []deck.Card{{Name: "h1", Value: 1, Suit: "h"}, {Name: "h10", Value: 10, Suit: "h"}},
			expectedResult: ResultOfIsStraight{IsStraight: true, HighestStraightCard: 14},
		},
	}

	for _, c := range cases {
		result, highestCard := IsStraight(c.board, c.pCards)
		fullResult := ResultOfIsStraight{IsStraight: result, HighestStraightCard: highestCard}

		if fullResult != c.expectedResult {
			t.Errorf("Expected %v, got %v", c.expectedResult.IsStraight, fullResult.IsStraight)
			t.Errorf("Expected %v, got %v", c.expectedResult.HighestStraightCard, fullResult.HighestStraightCard)
		}

	}
}

func TestPairToFullHouse(t *testing.T) {

	cases := []struct {
		board          []deck.Card
		pCards         []deck.Card
		expectedResult ResultOfPairToFullHousestruct
	}{
		{
			board:  []deck.Card{{Name: "h5", Value: 5, Suit: "h"}, {Name: "h6", Value: 6, Suit: "h"}, {Name: "h7", Value: 7, Suit: "h"}, {Name: "h13", Value: 13, Suit: "h"}, {Name: "h10", Value: 10, Suit: "h"}},
			pCards: []deck.Card{{Name: "h3", Value: 3, Suit: "h"}, {Name: "h2", Value: 2, Suit: "h"}},
			expectedResult: ResultOfPairToFullHousestruct{
				isPair:      false,
				score:       0,
				scoreString: "No Pair",
			},
		},
		{
			board:  []deck.Card{{Name: "h3", Value: 3, Suit: "h"}, {Name: "h6", Value: 6, Suit: "h"}, {Name: "h7", Value: 7, Suit: "h"}, {Name: "h13", Value: 13, Suit: "h"}, {Name: "h10", Value: 10, Suit: "h"}},
			pCards: []deck.Card{{Name: "h3", Value: 3, Suit: "h"}, {Name: "h2", Value: 2, Suit: "h"}},
			expectedResult: ResultOfPairToFullHousestruct{
				isPair:      true,
				score:       103,
				scoreString: "Pair",
			},
		},
		{
			board:  []deck.Card{{Name: "s3", Value: 3, Suit: "s"}, {Name: "h2", Value: 2, Suit: "h"}, {Name: "h7", Value: 7, Suit: "h"}, {Name: "h13", Value: 13, Suit: "h"}, {Name: "h10", Value: 10, Suit: "h"}},
			pCards: []deck.Card{{Name: "h3", Value: 3, Suit: "h"}, {Name: "c3", Value: 3, Suit: "c"}},
			expectedResult: ResultOfPairToFullHousestruct{
				isPair:      true,
				score:       203,
				scoreString: "Three of a Kind",
			},
		},
		{
			board:  []deck.Card{{Name: "s3", Value: 3, Suit: "s"}, {Name: "d3", Value: 3, Suit: "d"}, {Name: "s2", Value: 2, Suit: "s"}, {Name: "h13", Value: 13, Suit: "h"}, {Name: "h10", Value: 10, Suit: "h"}},
			pCards: []deck.Card{{Name: "h3", Value: 3, Suit: "h"}, {Name: "c3", Value: 3, Suit: "c"}},
			expectedResult: ResultOfPairToFullHousestruct{

				isPair:      true,
				score:       703,
				scoreString: "Four of a Kind",
			},
		},
		{
			board:  []deck.Card{{Name: "h3", Value: 3, Suit: "h"}, {Name: "h2", Value: 2, Suit: "h"}, {Name: "h7", Value: 7, Suit: "h"}, {Name: "h13", Value: 13, Suit: "h"}, {Name: "h10", Value: 10, Suit: "h"}},
			pCards: []deck.Card{{Name: "h3", Value: 3, Suit: "h"}, {Name: "h2", Value: 2, Suit: "h"}},
			expectedResult: ResultOfPairToFullHousestruct{
				isPair:      true,
				score:       155,
				scoreString: "Two Pair",
			},
		},

		{
			board:  []deck.Card{{Name: "s3", Value: 3, Suit: "s"}, {Name: "d2", Value: 2, Suit: "d"}, {Name: "s2", Value: 2, Suit: "s"}, {Name: "h13", Value: 13, Suit: "h"}, {Name: "h10", Value: 10, Suit: "h"}},
			pCards: []deck.Card{{Name: "h3", Value: 3, Suit: "h"}, {Name: "c3", Value: 3, Suit: "c"}},
			expectedResult: ResultOfPairToFullHousestruct{
				isPair:      true,
				score:       602,
				scoreString: "Full House",
			},
		},
		{
			board:  []deck.Card{{Name: "s2", Value: 2, Suit: "s"}, {Name: "d3", Value: 3, Suit: "d"}, {Name: "s3", Value: 3, Suit: "s"}, {Name: "h13", Value: 13, Suit: "h"}, {Name: "h10", Value: 10, Suit: "h"}},
			pCards: []deck.Card{{Name: "h2", Value: 2, Suit: "h"}, {Name: "c3", Value: 3, Suit: "c"}},
			expectedResult: ResultOfPairToFullHousestruct{
				isPair:      true,
				score:       603,
				scoreString: "Full House",
			},
		},
	}

	for _, c := range cases {
		isPair, score, string := PairToFullHouse(c.board, c.pCards)
		result := ResultOfPairToFullHousestruct{isPair, score, string}

		if result != c.expectedResult {
			t.Errorf("Expected %t, got %t", c.expectedResult.isPair, result.isPair)
			t.Errorf("Expected %d, got %d", c.expectedResult.score, result.score)
			t.Errorf("Expected %s, got %s", c.expectedResult.scoreString, result.scoreString)
		}
	}
}

type ResultOfPairToFullHousestruct struct {
	isPair      bool
	score       int
	scoreString string
	// value       int
}

type ResultOfIsFlush struct {
	isFlush          bool
	HighestFlushCard int
}

type ResultOfIsStraight struct {
	IsStraight          bool
	HighestStraightCard int
}
