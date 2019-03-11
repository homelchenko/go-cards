package main

func main() {
	cards := newDeck()
	prevCards := readFromFile("my_cards")
	prevCards.print()

	cards.shuffle()

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()

	cards.saveToFile("my_cards")
}
