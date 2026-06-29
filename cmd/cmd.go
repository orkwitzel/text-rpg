package cmd

import (
	"fmt"
	"rpg/cmd/commands"
	"rpg/cmd/utils"
	"rpg/internal/game"
	"strings"
)

func splitUserInputToArgs(input string) []string {
	return strings.Split(input, " ")
}

type command struct {
	name        string
	description string
	keywords    []string
	commandFunc func(*game.Game, []string) error
}

// GameInput is the main function for the game. It takes a game and a command and executes the command.
func GameLoop(g *game.Game) {
	fmt.Print("Enter command: ")
	args := splitUserInputToArgs(utils.InputString())
	command := commands.GetCommandFromInputArgs(args)
	if command == nil {
		fmt.Println("I don't know what do you mean")
		GameLoop(g)
	}
	err := command.CommandFunc(g, args)
	if err != nil {
		fmt.Println("Error:", err)
		GameLoop(g)
	}
}
