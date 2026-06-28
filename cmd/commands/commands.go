package commands

import (
	"fmt"
	"os"
	"rpg/cmd/utils"
	"rpg/internal/game"
	"slices"
)

type Command struct {
	name        string
	description string
	keywords    []string
	CommandFunc func(*game.Game, []string) error
}

func newCommand(name string, description string, keywords []string, commandFunc func(*game.Game, []string) error) Command {
	return Command{name: name, description: description, keywords: keywords, CommandFunc: commandFunc}
}

var baseCommands = []Command{
	newCommand("move", "Move the player in a direction", []string{"move", "walk", "go"}, moveCommand),
	newCommand("exit", "Exit the game", []string{"exit", "quit"}, exitCommand),
	newCommand("look", "Look around the current tile", []string{"look", "examine"}, lookCommand),
	newCommand("clear", "Clear the screen", []string{"clear"}, clearCommand),
}

var CommandsList = append(baseCommands, newCommand("help", "Show available commands", []string{"help"}, helpCommand))

// Returns the current command based on the input arguments.
func GetCommandFromInputArgs(args []string) *Command {
	for _, command := range CommandsList {
		if slices.Contains(command.keywords, args[0]) {
			return &command
		}
	}
	return nil
}

func lookCommand(g *game.Game, args []string) error {
	if len(args) > 1 {
		return fmt.Errorf("look command takes no arguments")
	}

	fmt.Println("You look around and see the following:")
	currentTile := g.GetPlayerTile()
	fmt.Println("Current tile:", currentTile.Name)
	fmt.Println("Enemies:", currentTile.Enemies)
	fmt.Println("Items:", currentTile.Items)

	fmt.Println("To the north you see a", g.World.GetTile(g.PlayerPositionX, g.PlayerPositionY+1).Name)
	fmt.Println("To the south you see a", g.World.GetTile(g.PlayerPositionX, g.PlayerPositionY-1).Name)
	fmt.Println("To the east you see a", g.World.GetTile(g.PlayerPositionX+1, g.PlayerPositionY).Name)
	fmt.Println("To the west you see a", g.World.GetTile(g.PlayerPositionX-1, g.PlayerPositionY).Name)

	return nil
}

func exitCommand(_ *game.Game, args []string) error {
	if len(args) > 1 {
		return fmt.Errorf("exit command takes no arguments")
	}

	fmt.Println("Exiting game")
	os.Exit(0)
	return nil
}

func clearCommand(_ *game.Game, args []string) error {
	if len(args) > 1 {
		return fmt.Errorf("clear command takes no arguments")
	}

	utils.ClearScreen()
	return nil
}

func moveCommand(g *game.Game, args []string) error {
	if len(args) > 2 {
		return fmt.Errorf("move command takes one argument")
	}

	var direction game.Direction
	switch args[1] {
	case "up":
		direction = game.DirectionNorth
	case "down":
		direction = game.DirectionSouth
	case "left":
		direction = game.DirectionWest
	case "right":
		direction = game.DirectionEast
	default:
		return fmt.Errorf("invalid direction")
	}

	g.MovePlayer(direction)
	fmt.Println("Player moved to", direction)
	fmt.Println("Player is now on", g.GetPlayerTile())
	return nil
}

func helpCommand(_ *game.Game, args []string) error {
	if len(args) > 1 {
		return fmt.Errorf("help command takes no arguments")
	}

	fmt.Println("Available commands:")
	for _, command := range baseCommands {
		fmt.Println(command.name, "-", command.description)
	}
	fmt.Println("help - Show available commands")
	return nil
}
