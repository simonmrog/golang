package main

import "fmt"

type deck []string

func newDeck() deck {
	var cards = deck{}

	var cardSuits = []string{"Spades",  "Diamonds", "Hearts", "Clubs"}
	var cardValues = []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}

	return cards
}

func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}
