package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()

	if len(deck) != 52 {
		t.Errorf("Incorrect number of cards in the deck. Deck length is %v", len(deck))
	}

	if deck[0] != "Ace of Spades" {
		t.Errorf("Incorrect card at the beginning of the deck. card is %v", deck[0])
	}

	if deck[len(deck) - 1] != "King of Diamonds" {
		t.Errorf("Incorrect card at the ending of the deck. card is %v", deck[len(deck) - 1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	const fileName = "_decktesting.txt"
	os.Remove(fileName)

	deck := newDeck()
	deck.saveToFile(fileName)

	loadedDeck := newDeckFromFile(fileName)

	if len(loadedDeck) != 52 {
		t.Errorf("Incorrect number of cards in the deck. Deck length is %v", len(deck))
	}

	os.Remove(fileName)
}