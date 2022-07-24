package analyser

import (
	"testing"

	"github.com/arctheowl/GoDeck/deck"
)

// TestIsStraight tests the IsHigh function
func TestCompareHands(t *testing.T) {

	cases := []struct {
		board          []deck.Card
		p1Cards        []deck.Card
		p2Cards        []deck.Card
		expectedResult int
	}{
		{
			board: []deck.Card{
				{Name: "h1", Value: 1, Suit: "h"},
				{Name: "h2", Value: 2, Suit: "h"},
				{Name: "h3", Value: 3, Suit: "h"},
				{Name: "h4", Value: 4, Suit: "h"},
				{Name: "h5", Value: 5, Suit: "h"},
			},
			p1Cards: []deck.Card{
				{Name: "s10", Value: 10, Suit: "s"}, {Name: "d2", Value: 2, Suit: "d"},
			},

			p2Cards: []deck.Card{
				{Name: "d10", Value: 10, Suit: "d"}, {Name: "s2", Value: 2, Suit: "s"},
			},
			expectedResult: 0,
		},
		{
			board: []deck.Card{
				{Name: "h1", Value: 1, Suit: "h"},
				{Name: "c2", Value: 2, Suit: "c"},
				{Name: "c3", Value: 3, Suit: "c"},
				{Name: "d4", Value: 4, Suit: "d"},
				{Name: "h5", Value: 5, Suit: "h"},
			},
			p1Cards: []deck.Card{
				{Name: "s10", Value: 10, Suit: "s"}, {Name: "d2", Value: 2, Suit: "d"},
			},

			p2Cards: []deck.Card{
				{Name: "d10", Value: 10, Suit: "d"}, {Name: "s2", Value: 2, Suit: "s"},
			},
			expectedResult: 0,
		},
		{
			board: []deck.Card{
				{Name: "h1", Value: 1, Suit: "h"},
				{Name: "c2", Value: 2, Suit: "c"},
				{Name: "c3", Value: 3, Suit: "c"},
				{Name: "d4", Value: 4, Suit: "d"},
				{Name: "h5", Value: 5, Suit: "h"},
			},
			p1Cards: []deck.Card{
				{Name: "s6", Value: 6, Suit: "s"}, {Name: "d2", Value: 2, Suit: "d"},
			},

			p2Cards: []deck.Card{
				{Name: "d10", Value: 10, Suit: "d"}, {Name: "s2", Value: 2, Suit: "s"},
			},
			expectedResult: 1,
		},
		{
			board: []deck.Card{
				{Name: "h12", Value: 12, Suit: "h"},
				{Name: "c2", Value: 2, Suit: "c"},
				{Name: "c3", Value: 3, Suit: "c"},
				{Name: "d4", Value: 4, Suit: "d"},
				{Name: "h5", Value: 5, Suit: "h"},
			},
			p1Cards: []deck.Card{
				{Name: "s11", Value: 11, Suit: "s"}, {Name: "d9", Value: 9, Suit: "d"},
			},

			p2Cards: []deck.Card{
				{Name: "d10", Value: 10, Suit: "d"}, {Name: "s9", Value: 9, Suit: "s"},
			},
			expectedResult: 1,
		},
		{
			board: []deck.Card{
				{Name: "h11", Value: 11, Suit: "h"},
				{Name: "c2", Value: 2, Suit: "c"},
				{Name: "c3", Value: 3, Suit: "c"},
				{Name: "d11", Value: 11, Suit: "d"},
				{Name: "h5", Value: 5, Suit: "h"},
			},
			p1Cards: []deck.Card{
				{Name: "s11", Value: 11, Suit: "s"}, {Name: "c11", Value: 11, Suit: "c"},
			},

			p2Cards: []deck.Card{
				{Name: "d10", Value: 10, Suit: "d"}, {Name: "s9", Value: 9, Suit: "s"},
			},
			expectedResult: 1,
		},
		{
			board: []deck.Card{
				{Name: "h1", Value: 1, Suit: "h"},
				{Name: "s8", Value: 8, Suit: "s"},
				{Name: "c7", Value: 7, Suit: "c"},
				{Name: "d10", Value: 10, Suit: "d"},
				{Name: "c1", Value: 5, Suit: "h"},
			},
			p1Cards: []deck.Card{
				{Name: "s13", Value: 13, Suit: "s"}, {Name: "c12", Value: 12, Suit: "c"},
			},

			p2Cards: []deck.Card{
				{Name: "3h", Value: 3, Suit: "h"}, {Name: "s2", Value: 2, Suit: "s"},
			},
			expectedResult: 1,
		},
		{
			board: []deck.Card{
				{Name: "h1", Value: 1, Suit: "h"},
				{Name: "s8", Value: 8, Suit: "s"},
				{Name: "c7", Value: 7, Suit: "c"},
				{Name: "d10", Value: 10, Suit: "d"},
				{Name: "c1", Value: 5, Suit: "h"},
			},
			p1Cards: []deck.Card{
				{Name: "3h", Value: 3, Suit: "h"}, {Name: "s2", Value: 2, Suit: "s"},
			},
			p2Cards: []deck.Card{
				{Name: "s13", Value: 13, Suit: "s"}, {Name: "c12", Value: 12, Suit: "c"},
			},
			expectedResult: 2,
		},
		{
			board: []deck.Card{
				{Name: "h1", Value: 1, Suit: "h"},
				{Name: "s8", Value: 8, Suit: "s"},
				{Name: "c7", Value: 7, Suit: "c"},
				{Name: "d10", Value: 10, Suit: "d"},
				{Name: "c1", Value: 5, Suit: "h"},
			},
			p1Cards: []deck.Card{
				{Name: "h10", Value: 10, Suit: "h"}, {Name: "s2", Value: 2, Suit: "s"},
			},
			p2Cards: []deck.Card{
				{Name: "s10", Value: 10, Suit: "s"}, {Name: "c12", Value: 12, Suit: "c"},
			},
			expectedResult: 2,
		},
		{
			board: []deck.Card{
				{Name: "h1", Value: 1, Suit: "h"},
				{Name: "s8", Value: 8, Suit: "s"},
				{Name: "c7", Value: 7, Suit: "c"},
				{Name: "d10", Value: 10, Suit: "d"},
				{Name: "c1", Value: 5, Suit: "h"},
			},
			p1Cards: []deck.Card{
				{Name: "s10", Value: 10, Suit: "s"}, {Name: "c12", Value: 12, Suit: "c"},
			},
			p2Cards: []deck.Card{
				{Name: "h10", Value: 10, Suit: "h"}, {Name: "s2", Value: 2, Suit: "s"},
			},
			expectedResult: 1,
		},
	}

	for _, c := range cases {
		result := CompareHands(c.board, c.p1Cards, c.p2Cards)
		if result != c.expectedResult {
			t.Errorf("Expected %d, got %d", c.expectedResult, result)
		}
	}
}

func TestAnalyse(t *testing.T) {
	cases := []struct {
		board          []deck.Card
		p1Cards        []deck.Card
		expectedResult ResultOfAnalyse
	}{
		{
			board: []deck.Card{
				{Name: "h1", Value: 1, Suit: "h"},
				{Name: "h2", Value: 2, Suit: "h"},
				{Name: "h3", Value: 3, Suit: "h"},
				{Name: "h4", Value: 4, Suit: "h"},
				{Name: "h6", Value: 6, Suit: "h"},
			},
			p1Cards: []deck.Card{
				{Name: "s10", Value: 10, Suit: "s"}, {Name: "d2", Value: 2, Suit: "d"},
			},
			expectedResult: ResultOfAnalyse{500, "Flush"},
		},
		{
			board: []deck.Card{
				{Name: "h1", Value: 1, Suit: "h"},
				{Name: "h2", Value: 2, Suit: "h"},
				{Name: "h3", Value: 3, Suit: "h"},
				{Name: "h4", Value: 4, Suit: "h"},
				{Name: "h5", Value: 5, Suit: "h"},
			},
			p1Cards: []deck.Card{
				{Name: "s10", Value: 10, Suit: "s"}, {Name: "d2", Value: 2, Suit: "d"},
			},
			expectedResult: ResultOfAnalyse{905, "Straight Flush"},
		},
		{
			board: []deck.Card{
				{Name: "h10", Value: 10, Suit: "h"},
				{Name: "h11", Value: 11, Suit: "h"},
				{Name: "h12", Value: 12, Suit: "h"},
				{Name: "d1", Value: 1, Suit: "d"},
				{Name: "c1", Value: 1, Suit: "c"},
			},
			p1Cards: []deck.Card{
				{Name: "h1", Value: 1, Suit: "h"}, {Name: "h13", Value: 13, Suit: "h"},
			},
			expectedResult: ResultOfAnalyse{1000, "Royal Flush"},
		},
	}

	for _, c := range cases {
		score, scoreString := Analyse(c.board, c.p1Cards)
		result := ResultOfAnalyse{score, scoreString}
		if result != c.expectedResult {
			t.Errorf("Expected %d, got %d", c.expectedResult.score, result.score)
			t.Errorf("Expected %s, got %s", c.expectedResult.scoreString, result.scoreString)
		}
	}

}

type ResultOfAnalyse struct {
	score       int
	scoreString string
}
