package main

func main() {
	// var card string = "Ace of spades"
	// cards := newDeck()

	// cards.saveTotFile(("my_cards"))

	// fmt.Println(cards.toString())

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	// type conversion example
	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))
	// Output: [72 105 32 116 104 101 114 101 33]

	// cards := newDeckFromFile("my_cards")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
