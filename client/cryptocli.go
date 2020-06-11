package client

import (
	"fmt"
)

type deck []string

func NewDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades","Diamonds","Hearts","Clubs"}
	cardValues := []string{"Ace","Two","Three","Four"}

	for _, suit := range cardSuits {
		for _, values := range cardValues {
			cards= append(cards, values+" of "+suit)
		}
	}

	return cards
}

func(d deck) Print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}