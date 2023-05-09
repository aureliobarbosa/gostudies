package main

import "fmt"

func main() {
	fmt.Println("Initial test!")
	deck := NewDeck()
	deck.shuffle()
	deck.print()	
}