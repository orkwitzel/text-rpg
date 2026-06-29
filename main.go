package main

import (
	"rpg/cmd"
	savefiles "rpg/internal/saveFiles"
)

func main() {
	game := cmd.SaveFilesMenu()

	for {
		cmd.GameLoop(&game)
		game.WorldInteraction()
		savefiles.SaveGame(game)
	}
}
