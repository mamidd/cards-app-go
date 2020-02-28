package main

import "fmt"

func main() {
	cards := []string{"a", "b"}

	for card := range cards {
		fmt.Println(card)
	}
}
