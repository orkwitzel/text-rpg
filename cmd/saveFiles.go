package cmd

import (
	"fmt"
	"os"
	"rpg/cmd/utils"
	"rpg/internal/game"
	"rpg/internal/game/player"
	"rpg/internal/game/world"
	savefiles "rpg/internal/saveFiles"
	worldloader "rpg/internal/world-loader"
)

func SaveFilesMenu() game.Game {
	fmt.Println("1. Load saved game")
	fmt.Println("2. Create new world")
	fmt.Println("3. Load world from directory")
	fmt.Print("Choose an option: ")

	switch utils.InputString() {
	case "1", "load":
		return loadSavedGame()
	case "2", "new", "create":
		return createNewGame()
	case "3", "load-world":
		return loadWorldFromDirectory()
	default:
		fmt.Println("Invalid choice. Enter 1, 2, or 3.")
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
	choice := utils.InputInt()
	if choice < 1 || choice > len(saveFiles) {
		fmt.Println("Invalid choice.")
		return loadSavedGame()
	}

	selected := saveFiles[choice-1]
	fmt.Printf("Loaded %s.\n", selected.Game.World.Name)
	return selected.Game
}

func loadWorldFromDirectory() game.Game {
	fmt.Print("World directory: ")
	worldDir := utils.InputString()
	if _, err := os.Stat(worldDir); os.IsNotExist(err) {
		utils.ClearScreen()
		fmt.Println("World directory does not exist.")
		return SaveFilesMenu()
	}
	return worldloader.LoadWorld(worldDir)
}

func createNewGame() game.Game {
	fmt.Print("World name: ")
	worldName := utils.InputString()
	fmt.Print("World width: ")
	width := utils.InputInt()
	fmt.Print("World height: ")
	height := utils.InputInt()
	worldMap := world.New(worldName, width, height)

	fmt.Print("Player name: ")
	playerName := utils.InputString()
	player := player.New(playerName)

	return game.New(player, worldMap)
}
