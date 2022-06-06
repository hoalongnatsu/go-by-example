package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var players []string

	for {
		var player string
		fmt.Print("Please enter player's name or enter 'start' to starting game: ")
		fmt.Scan(&player)
		players = append(players, player)

		if player == "start" {
			break
		}
	}

	fmt.Printf("Have %d players: ", len(players))
	for _, v := range players {
		fmt.Printf("%s ", v)
	}

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	number := r.Intn(len(players))

	fmt.Printf("\nRandom player is %s\n", players[number])
}
