package main

import "fmt"

type CardNumber int64
type Suit string

type Card struct {
	value CardNumber
	suit  Suit
}

var cardNumberStrings = []string{"A", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
var suitStrings = []string{"♥", "♠", "♦", "♣"}

type deck []Card

func (c Card) toString() string {
	return cardNumberStrings[c.value] + string(c.suit)
}

func NewDeck() {
	// maxNumber = len(cardNumberStrings)

	// var myDeck deck
	// var thisCard Card

	for _, suitString := range suitStrings {
		for _, cardNumberString := range cardNumberStrings {
			fmt.Printf("%v%v ", cardNumberString, suitString)
		}
		fmt.Println()
	}
}
