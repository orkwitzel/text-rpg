package tiles

import (
	"math/rand"

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
	ID   string   `json:"id"`
	Name string   `json:"name"`
	Type TileType `json:"type"`
}

func New(name string, tileType TileType) Tile {
	return Tile{
		ID:   uuid.NewString(),
		Name: name,
		Type: tileType,
	}
}

func NewRandom() Tile {
	types := []TileType{TileTypeGrass, TileTypeWater, TileTypeSand, TileTypeRock}
	t := types[rand.Intn(len(types))]
	return New(string(t), t)
}
