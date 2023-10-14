package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new deck
type deck []string

func newDeck() deck {
	cardSuits := [4]string{"Spades", "Hearts", "Clubs", "Diamonds"}
	cardValues := [13]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	var newDeck deck

	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			newDeck = append(newDeck, cardValue + " of " + cardSuit)
		}
	}
	return newDeck
}

func newDeckFromFile(fileName string) deck {
	bs, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), ", "))
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//--------Deck receiver methods--------

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) shuffle() deck {
	rand.New(rand.NewSource(time.Now().Local().UnixNano()))
	rand.Shuffle(len(d), func(i, j int) {d[i], d[j] = d[j], d[i]})
	return d
}

func (d deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(d.toString()), 0666)
}

func (d deck) toString() string {
	return strings.Join([]string(d), ", ")
}