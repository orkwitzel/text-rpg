# text-rpg

A terminal-based RPG written in Go. Explore a randomly generated tile world, move your player across the map, and persist your progress with a built-in save system.

## Features

- Randomly generated world maps with four tile types: grass, water, sand, and rock
- Player movement in four directions (up, down, left, right)
- JSON-based save/load system — games are stored in `~/rpg/saveFiles/`
- Multiple save slots; pick up where you left off

## Requirements

- Go 1.21+

## Installation

```bash
git clone https://github.com/orkwitzel/text-rpg.git
cd text-rpg
go build -o rpg .
./rpg
```

Or run directly without building:

```bash
go run .
```

## Usage

On startup, choose to load an existing save or create a new world:

```
1. Load saved game
2. Create new world
Choose an option:
```

When creating a new world you will be prompted for a world name, dimensions, and a player name.

In-game commands:

| Command         | Description              |
|-----------------|--------------------------|
| `move up`       | Move the player north    |
| `move down`     | Move the player south    |
| `move left`     | Move the player west     |
| `move right`    | Move the player east     |
| `exit`          | Save and quit            |

The game auto-saves after every command.

## Project Structure

```
.
├── main.go                        # Entry point
├── cmd/
│   ├── cmd.go                     # Input helpers (string, int, float, bool)
│   ├── gameCmd.go                 # In-game command handling
│   └── saveFiles.go               # Save/load menu
└── internal/
    ├── game/
    │   ├── game.go                # Game state and player movement
    │   ├── player/
    │   │   └── player.go          # Player model
    │   └── world/
    │       ├── world.go           # World map model
    │       └── tiles/
    │           └── tile.go        # Tile types and random generation
    └── saveFiles/
        └── saveFiles.go           # JSON persistence layer
```

## Save Files

Save files are stored at `~/rpg/saveFiles/<game-id>.json`. Each file contains the full world map, player state, and timestamps.
