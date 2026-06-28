package enemy

import (
	"rpg/internal/game/stats"

	"github.com/google/uuid"
)

type Enemy struct {
	ID     string      `json:"id"`
	Name   string      `json:"name"`
	Stats  stats.Stats `json:"stats"`
	Level  int         `json:"level"`
	Health int         `json:"health"`
	Dead   bool        `json:"dead"`
}

func NewEnemy(name string, stats stats.Stats, health int) *Enemy {
	return &Enemy{ID: uuid.New().String(),
		Name: name, Stats: stats, Health: health,
		Dead: false}
}

func (e *Enemy) IsDead() bool {
	return e.Dead
}

func (e *Enemy) TakeDamage(damage int) {
	e.Health -= damage
	if e.Health <= 0 {
		e.Dead = true
	}
}
