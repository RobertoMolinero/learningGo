package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	expectedLength := 52
	expectedFirstCard := "Ace of Spades"
	expectedLastCard := "King of Clubs"

	d := newDeck()

	lengthOfDeck := len(d)

	if lengthOfDeck != expectedLength {
		t.Errorf("Expected Deck Length of %v, but got %v", expectedLength, lengthOfDeck)
	}

	firstCard := d[0]

	if firstCard != expectedFirstCard {
		t.Errorf("Expected First Card is %v, but got %v", expectedFirstCard, firstCard)
	}

	lastCard := d[expectedLength-1]

	if lastCard != expectedLastCard {
		t.Errorf("Expected First Card is %v, but got %v", expectedLastCard, lastCard)
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {

	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loaded := newDeckFromFile("_decktesting")

	expectedLength := 52
	lengthOfDeck := len(loaded)

	if lengthOfDeck != expectedLength {
		t.Errorf("Expected Deck Length of %v, but got %v", expectedLength, lengthOfDeck)
	}

	os.Remove("_decktesting")
}