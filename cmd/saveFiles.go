package cmd

import (
	"fmt"
	"rpg/internal/game"
	"rpg/internal/game/player"
	"rpg/internal/game/world"
	savefiles "rpg/internal/saveFiles"
)

func SaveFilesMenu() game.Game {
	fmt.Println("1. Load saved game")
	fmt.Println("2. Create new world")
	fmt.Print("Choose an option: ")

	switch InputString() {
	case "1", "load":
		return loadSavedGame()
	case "2", "new", "create":
		return createNewGame()
	default:
		fmt.Println("Invalid choice. Enter 1 or 2.")
		return SaveFilesMenu()
	}
}

func loadSavedGame() game.Game {
	saveFiles := savefiles.ListSaveFiles()
	if len(saveFiles) == 0 {
		fmt.Println("No save files found.")
		return SaveFilesMenu()
	}

	fmt.Println("Available save files:")
	for i, saveFile := range saveFiles {
		fmt.Printf(
			"%d. %s (player: %s)\n",
			i+1,
			saveFile.Game.World.Name,
			saveFile.Game.Player.Name,
		)
	}

	fmt.Print("Choose a save file: ")
	choice := InputInt()
	if choice < 1 || choice > len(saveFiles) {
		fmt.Println("Invalid choice.")
		return loadSavedGame()
	}

	selected := saveFiles[choice-1]
	fmt.Printf("Loaded %s.\n", selected.Game.World.Name)
	return selected.Game
}

func createNewGame() game.Game {
	fmt.Print("World name: ")
	worldName := InputString()
	fmt.Print("World width: ")
	width := InputInt()
	fmt.Print("World height: ")
	height := InputInt()
	worldMap := world.New(worldName, width, height)

	fmt.Print("Player name: ")
	playerName := InputString()
	player := player.New(playerName)

	return game.New(player, worldMap)
}
