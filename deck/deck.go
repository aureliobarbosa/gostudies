package main

import (
	"fmt"
	"math/rand"
)

type CardNumber int64  // TODO: Validate CardNumber
type Suit string	// TODO: Validate Suit

type Card struct {
	number CardNumber
	numberString string
	suit  Suit
	stamp string
	value int64	
}

var cardNumberStrings = []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
var suitStrings = []string{"♥", "♠", "♦", "♣"}

type Deck []Card

func (deck Deck) print() {
	for _, card := range deck {
		if len(card.stamp) != 0 {
			fmt.Printf("%v ",card.stamp)
		} else {
			fmt.Printf("%v%v ",card.numberString, string(card.suit))
		}
	}
	fmt.Println()
}

func (deck Deck) shuffle() {	
	var j int

	n := len(deck)
	for i := 0; i < n; i++ {		
		j = rand.Intn(n)
		deck[i], deck[j] = deck[j], deck[i]
	}
}

func NewDeck() Deck {
	deck := Deck{}
	card := new(Card)

	for _, suitString := range suitStrings {
		for i, cardNumberString := range cardNumberStrings {
			card.number = CardNumber(i + 1)
			card.numberString = cardNumberString
			card.suit = Suit(suitString)
			card.value  = 0
			card.stamp = cardNumberString + suitString
			deck = append(deck,*card)
		}
	}
	return deck
}