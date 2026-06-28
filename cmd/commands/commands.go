package commands

import (
	"fmt"
	"os"
	"rpg/cmd/utils"
	"rpg/internal/game"
	"rpg/internal/game/battlesys"
	"rpg/internal/game/npcs/enemy"
	"rpg/internal/game/world/tiles"
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
	newCommand("attack", "Attack an enemy", []string{"attack", "fight", "kill"}, playerAttackCommand),
}

var CommandsList = append(baseCommands, newCommand("help", "Show available commands", []string{"help"}, helpCommand))

// Returns the current command based on the input arguments.
func GetCommandFromInputArgs(args []string) *Command {
	args = utils.RemoveLinkingWordsFromArgs(args)
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

	lookDirection(g, "north", g.PlayerPositionX, g.PlayerPositionY+1)
	lookDirection(g, "south", g.PlayerPositionX, g.PlayerPositionY-1)
	lookDirection(g, "east", g.PlayerPositionX+1, g.PlayerPositionY)
	lookDirection(g, "west", g.PlayerPositionX-1, g.PlayerPositionY)

	return nil
}

func lookDirection(g *game.Game, direction string, x, y int) {
	if !g.World.InBounds(x, y) {
		fmt.Println("To the", direction, "you see the edge of the world.")
		return
	}
	fmt.Println("To the", direction, "you see a", g.World.GetTile(x, y).Name)
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

func playerAttackCommand(g *game.Game, args []string) error {
	currentTile := g.World.TileAt(g.PlayerPositionX, g.PlayerPositionY)

	if len(currentTile.Enemies) == 0 {
		return fmt.Errorf("no enemies on this tile")
	}

	var targetEnemy *enemy.Enemy
	if len(args) > 1 {
		targetEnemy = tiles.LocateEnemyBasedOnName(args[1], currentTile)
		if targetEnemy == nil {
			return fmt.Errorf("enemy not found")
		}
	} else if len(currentTile.Enemies) > 1 {
		return fmt.Errorf("multiple enemies on this tile, please specify which one to attack")
	} else {
		targetEnemy = &currentTile.Enemies[0]
	}

	if targetEnemy.IsDead() {
		return fmt.Errorf("enemy is already dead")
	}

	damage := battlesys.CalculateDamageToEnemy(targetEnemy, &g.Player)
	if damage == 0 {
		fmt.Println("You missed", targetEnemy.Name)
		return nil
	}

	targetEnemy.TakeDamage(damage)
	fmt.Printf("You hit %s for %d damage\n", targetEnemy.Name, damage)
	if targetEnemy.IsDead() {
		fmt.Println(targetEnemy.Name, "has been defeated")
	} else {
		fmt.Println(targetEnemy.Name, "has", targetEnemy.Health, "health left")
	}
	return nil
}
