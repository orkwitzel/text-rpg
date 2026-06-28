package worldloader

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"rpg/internal/game"
	"rpg/internal/game/item"
	"rpg/internal/game/npcs/enemy"
	"rpg/internal/game/player"
	"rpg/internal/game/world"
	"rpg/internal/game/world/tiles"
)

func LoadWorld(worldDir string) game.Game {
	metadata := entireWorldMetadata{
		WorldMetadata: loadGenericMetadata[worldMetadata](worldDir, "metadata.json"),
		Enemies:       loadGenericMetadata[[]enemyMetadata](worldDir, "enemies.json"),
		Items:         loadGenericMetadata[[]itemMetadata](worldDir, "items.json"),
		Tiles:         loadGenericMetadata[[]tileMetadata](worldDir, "tiles.json"),
		Player:        loadGenericMetadata[playerMetadata](worldDir, "player.json"),
	}
	return entireWorldMetadataToGame(metadata)
}

type entireWorldMetadata struct {
	WorldMetadata worldMetadata   `json:"world_metadata"`
	Enemies       []enemyMetadata `json:"enemies"`
	Items         []itemMetadata  `json:"items"`
	Tiles         []tileMetadata  `json:"tiles"`
	Player        playerMetadata  `json:"player"`
}

type worldMetadata struct {
	Name                   string `json:"name"`
	Version                int    `json:"version"`
	TileNextRowEveryNTiles int    `json:"tilesNextRowEveryNTiles"`
}

type enemyMetadata struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Health int    `json:"health"`
	Damage int    `json:"damage"`
	Speed  int    `json:"speed"`
	Level  int    `json:"level"`
}

type itemMetadata struct {
	ID              string         `json:"id"`
	Name            string         `json:"name"`
	Description     string         `json:"description"`
	Value           int            `json:"value"`
	Weight          int            `json:"weight"`
	CanHaveMultiple bool           `json:"can_have_multiple"`
	Quantity        int            `json:"quantity"`
	ItemType        string         `json:"item_type"`
	Effects         effectMetadata `json:"effects"`
	RequiredLevel   int            `json:"required_level"`
}

type effectMetadata struct {
	Health  int `json:"health"`
	Energy  int `json:"energy"`
	Damage  int `json:"damage"`
	Defense int `json:"defense"`
	Speed   int `json:"speed"`
	Level   int `json:"level"`
}

type tileMetadata struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Type        tiles.TileType `json:"type"`
	Enemies     []string       `json:"enemies"`
	Items       []string       `json:"items"`
	Description string         `json:"description"`
}

type playerMetadata struct {
	Name                  string `json:"name"`
	Health                int    `json:"health"`
	Energy                int    `json:"energy"`
	Level                 int    `json:"level"`
	Experience            int    `json:"experience"`
	ExperienceToNextLevel int    `json:"experience_to_next_level"`
	Gold                  int    `json:"gold"`
	InitialPositionX      int    `json:"initial_position_x"`
	InitialPositionY      int    `json:"initial_position_y"`
}

func loadGenericMetadata[T any](worldDir string, filename string) T {
	filePath := filepath.Join(worldDir, filename)
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read world %s: %v", filename, err)
	}

	var data T
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		log.Fatalf("Failed to unmarshal world %s: %v", filename, err)
	}
	return data
}

func entireWorldMetadataToGame(metadata entireWorldMetadata) game.Game {
	p := playerMetadataToPlayer(metadata.Player)
	tilesArr := []tiles.Tile{}
	for _, tile := range metadata.Tiles {
		tilesArr = append(tilesArr, tileMetadataToTile(tile))
	}
	enemies := []enemy.Enemy{}
	for _, enemy := range metadata.Enemies {
		enemies = append(enemies, enemyMetadataToEnemy(enemy))
	}
	items := []item.Item{}
	for _, item := range metadata.Items {
		items = append(items, itemMetadataToItem(item))
	}

	tilesMap := make([][]tiles.Tile, 0)
	for i, row := range tilesArr {
		if i%metadata.WorldMetadata.TileNextRowEveryNTiles == 0 {
			tilesMap = append(tilesMap, make([]tiles.Tile, 0))
		}
		tilesMap[len(tilesMap)-1] = append(tilesMap[len(tilesMap)-1], row)
	}

	w := world.New(metadata.WorldMetadata.Name, len(tilesMap[0]), len(tilesMap))
	for y, row := range tilesMap {
		for x, tile := range row {
			w.SetTile(x, y, tile)
		}
	}

	g := game.New(p, w)
	g.PlayerPositionX = metadata.Player.InitialPositionX
	g.PlayerPositionY = metadata.Player.InitialPositionY
	return g
}

func playerMetadataToPlayer(metadata playerMetadata) player.Player {
	p := player.New(metadata.Name)
	p.Health = metadata.Health
	p.Energy = metadata.Energy
	p.Level = metadata.Level
	p.Experience = metadata.Experience
	p.ExperienceToNextLevel = metadata.ExperienceToNextLevel
	p.Gold = metadata.Gold
	return p
}

func tileMetadataToTile(metadata tileMetadata) tiles.Tile {
	return tiles.New(metadata.Name, metadata.Type, make([]enemy.Enemy, 0), make([]item.Item, 0))
}

func enemyMetadataToEnemy(metadata enemyMetadata) enemy.Enemy {
	return *enemy.NewEnemy(metadata.Name, metadata.Health, metadata.Damage, metadata.Speed, metadata.Level)
}

func itemMetadataToItem(metadata itemMetadata) item.Item {
	return *item.NewItem(metadata.Name, metadata.Description, metadata.Value, metadata.Weight, metadata.CanHaveMultiple, metadata.Quantity)
}
