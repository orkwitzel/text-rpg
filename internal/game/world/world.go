package world

import "rpg/internal/game/world/tiles"

type WorldMap struct {
	Name   string         `json:"name"`
	Width  int            `json:"width"`
	Height int            `json:"height"`
	Tiles  [][]tiles.Tile `json:"tiles"`
}

func New(name string, width int, height int) WorldMap {
	tilesMap := make([][]tiles.Tile, 0)

	for range height {
		row := make([]tiles.Tile, 0)
		for range width {
			row = append(row, tiles.NewRandom())
		}
		tilesMap = append(tilesMap, row)
	}

	return WorldMap{
		Name:   name,
		Width:  width,
		Height: height,
		Tiles:  tilesMap,
	}
}

func (m *WorldMap) InBounds(x int, y int) bool {
	return x >= 0 && x < m.Width && y >= 0 && y < m.Height
}

func (m *WorldMap) GetTile(x int, y int) tiles.Tile {
	return m.Tiles[y][x]
}

func (m *WorldMap) TileAt(x int, y int) *tiles.Tile {
	return &m.Tiles[y][x]
}

func (m *WorldMap) SetTile(x int, y int, tile tiles.Tile) {
	m.Tiles[y][x] = tile
}
