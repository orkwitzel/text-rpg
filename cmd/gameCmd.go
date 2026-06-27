package cmd

import (
	"fmt"
	"os"
	"rpg/internal/game"
	"strings"
)

// GameInput is the main function for the game. It takes a game and a command and executes the command.
func GameInput(g *game.Game) {
	fmt.Print("Enter command: ")
	command := InputString()
	if strings.HasPrefix(command, "move") {
		direction := moveCommandToDirection(command)
		if direction == nil {
			fmt.Println("Invalid direction. Options are: up, down, left, right")
			return
		}
		g.MovePlayer(*direction)
		fmt.Println("Player moved to", *direction)
		fmt.Println("Player is now on", g.GetPlayerTile())
	} else if strings.HasPrefix(command, "exit") {
		fmt.Println("Exiting game")
		os.Exit(0)
	}
}

// DirectionInput is a helper function for the game. It takes a game and a direction and returns the direction.
func moveCommandToDirection(moveCommand string) *game.Direction {
	var direction game.Direction
	directionString := removePrefix(moveCommand, "move")

	switch directionString {
	case "up":
		direction = game.DirectionUp
	case "down":
		direction = game.DirectionDown
	case "left":
		direction = game.DirectionLeft
	case "right":
		direction = game.DirectionRight
	default:
		return nil
	}

	return &direction
}

func removePrefix(s string, prefix string) string {
	return strings.TrimSpace(strings.TrimPrefix(s, prefix))
}
