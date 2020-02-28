package main

func main() {
	cards := newDeck()

	hand, remaininDeck := deal(cards, 5)

	hand.print()
	remaininDeck.print()
}
