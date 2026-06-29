package main

import (
	"fmt"
	"rpg/cmd"
	savefiles "rpg/internal/saveFiles"
)

func main() {
	game := cmd.SaveFilesMenu()

	for !game.Player.IsDead() {
		savefiles.SaveGame(game)
		cmd.GameLoop(&game)
		if game.Player.IsDead() {
			break
		}
		game.WorldInteraction()
	}

	fmt.Println("You are dead")
	fmt.Println("Game over")
}
