package tiles

import (
	"math/rand"
	"rpg/internal/game/item"
	"rpg/internal/game/npcs/enemy"

	"github.com/google/uuid"
)

type TileType string

const (
	TileTypeGrass TileType = "grass"
	TileTypeWater TileType = "water"
	TileTypeSand  TileType = "sand"
	TileTypeRock  TileType = "rock"
)

type Tile struct {
	ID      string        `json:"id"`
	Name    string        `json:"name"`
	Type    TileType      `json:"type"`
	Enemies []enemy.Enemy `json:"enemies"`
	Items   []item.Item   `json:"items"`
}

func New(name string, tileType TileType, enemies []enemy.Enemy, items []item.Item) Tile {
	return Tile{
		ID:      uuid.NewString(),
		Name:    name,
		Type:    tileType,
		Enemies: enemies,
		Items:   items,
	}
}

func NewRandom() Tile {
	types := []TileType{TileTypeGrass, TileTypeWater, TileTypeSand, TileTypeRock}
	t := types[rand.Intn(len(types))]
	return New(string(t), t, make([]enemy.Enemy, 0), make([]item.Item, 0))
}
