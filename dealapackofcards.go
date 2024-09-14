package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	players := 4
	cardsPerPlayer := len(deck) / players

	for i := 0; i < players; i++ {
		playerCards := deck[i*cardsPerPlayer : (i+1)*cardsPerPlayer]
		fmt.Printf("Player %d: ", i+1)
		for j, card := range playerCards {
			fmt.Print(card)
			if j != len(playerCards)-1 {
				fmt.Print(", ")
			}
		}
		fmt.Println()
	}
}
