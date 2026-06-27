package enemy

import "github.com/google/uuid"

type Enemy struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Health int    `json:"health"`
	Damage int    `json:"damage"`
	Speed  int    `json:"speed"`
	Level  int    `json:"level"`
}

func NewEnemy(name string, health int, damage int, speed int, level int) *Enemy {
	return &Enemy{ID: uuid.New().String(), Name: name, Health: health, Damage: damage, Speed: speed, Level: level}
}
