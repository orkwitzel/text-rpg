package main

import (
	"rpg/cmd"
	savefiles "rpg/internal/saveFiles"
)

func main() {
	game := cmd.SaveFilesMenu()

	for {
		cmd.GameInput(&game)
		savefiles.SaveGame(game)
	}
}
