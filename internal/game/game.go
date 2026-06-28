package game

import (
	"rpg/internal/game/player"
	"rpg/internal/game/world"
	"rpg/internal/game/world/tiles"

	"github.com/google/uuid"
)

type Game struct {
	ID              string         `json:"id"`
	World           world.WorldMap `json:"world"`
	Player          player.Player  `json:"player"`
	PlayerPositionX int            `json:"player_position_x"`
	PlayerPositionY int            `json:"player_position_y"`
}

func New(p player.Player, w world.WorldMap) Game {
	return Game{
		ID:              uuid.NewString(),
		World:           w,
		Player:          p,
		PlayerPositionX: w.Width / 2,
		PlayerPositionY: w.Height / 2,
	}
}

type Direction string

const (
	DirectionNorth Direction = "north"
	DirectionSouth Direction = "south"
	DirectionEast  Direction = "east"
	DirectionWest  Direction = "west"
)

func (g *Game) MovePlayer(direction Direction) {
	switch direction {
	case DirectionNorth:
		g.PlayerPositionY--
	case DirectionSouth:
		g.PlayerPositionY++
	case DirectionEast:
		g.PlayerPositionX--
	case DirectionWest:
		g.PlayerPositionX++
	}
}

func (g *Game) GetPlayerTile() tiles.Tile {
	return g.World.GetTile(g.PlayerPositionX, g.PlayerPositionY)
}
