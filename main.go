package main

func main() {
	// 1
	// cards := newDeck()
	// cards = append(cards, "Six of Spades")

	// hand, remainingDeck := deal(cards, 5)

	// hand.print()
	// fmt.Println("---")
	// remainingDeck.print()

	// 2
	// cards := newDeck()

	// cards.saveToFile("mycard.txt")

	// 3
	// cards := newDeckFromFile("mycard.txt")
	// cards.print()

	// 4
	cards := newDeck()
	cards.shuffle()
	cards.print()

}
