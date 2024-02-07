package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	Deckarr := deck
	Player := 1

	for i := 0; i < 12; i += 3 {
		fmt.Printf("Player %v: %v, %v, %v\n", Player, Deckarr[i], Deckarr[i+1], Deckarr[i+2])
		Player++
	}
}
