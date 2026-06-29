package tiles

import (
	"fmt"
	"math/rand"
	"rpg/internal/game/item"
	"rpg/internal/game/npcs/enemy"
	"rpg/internal/helper"
	"strings"

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

func LocateEnemyBasedOnName(name string, t *Tile) *enemy.Enemy {
	formatted := strings.TrimSpace(strings.ToLower(name))
	if formatted == "" {
		return nil
	}
	for i := range t.Enemies {
		if strings.ToLower(t.Enemies[i].Name) == formatted {
			return &t.Enemies[i]
		}
	}
	return nil
}

func LocateItemBasedOfName(name string, t *Tile) *item.Item {
	formatted := strings.TrimSpace(strings.ToLower(name))
	if formatted == "" {
		return nil
	}
	for i := range t.Items {
		if strings.ToLower(t.Items[i].Name) == formatted {
			return &t.Items[i]
		}
	}
	return nil
}

func RemoveItemFromTile(i item.Item, t *Tile) error {
	for x, v := range t.Items {
		if v.ID == i.ID {
			t.Items = helper.RemoveElementFromArray(t.Items, x)
			return nil
		}
	}

	return fmt.Errorf("Item doesn't exist in tile")
}
