package main

import "fmt"

func DealAPackOfCards(deck []int) {
	Deckarr := deck
	Player := 1

	for i := 0; i < 12; i += 3 {
		fmt.Printf("Player %v: %v,%v,%v\n", Player, Deckarr[i], Deckarr[i+1], Deckarr[i+2])
		Player++
	}
}

func main() {
	deck := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	DealAPackOfCards(deck)
}
