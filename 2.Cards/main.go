package main

import "fmt"


var card = "Ace of Spades"

func main() {
	card := newCard()
	fmt.Println(card)
}

func newCard() string {
	return "Ace of Diamonds"
}
