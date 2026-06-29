package player

import (
	"rpg/internal/game/item"
	"rpg/internal/game/stats"
)

type Player struct {
	Name                  string      `json:"name"`
	Stats                 stats.Stats `json:"stats"`
	Health                int         `json:"health"`
	MaxHealth             int         `json:"max_health"`
	Energy                int         `json:"energy"`
	MaxEnergy             int         `json:"max_energy"`
	Level                 int         `json:"level"`
	Experience            int         `json:"experience"`
	ExperienceToNextLevel int         `json:"experience_to_next_level"`
	Gold                  int         `json:"gold"`
	EquippedWeapon        *item.Item  `json:"equipped_weapon"`
	Inventory             []item.Item `json:"inventory"`
}

func New(name string) Player {
	return Player{
		Name:                  name,
		Stats:                 stats.NewStats(10, 10, 10, 10, 10),
		Health:                100,
		MaxHealth:             100,
		Energy:                100,
		MaxEnergy:             100,
		Level:                 1,
		Experience:            0,
		ExperienceToNextLevel: 100,
		Gold:                  0,
		EquippedWeapon:        nil,
		Inventory:             make([]item.Item, 0),
	}
}

func (p *Player) IsDead() bool {
	return p.Health <= 0
}

func (p *Player) TakeDamage(damage int) {
	p.Health -= damage
}

func (p *Player) TakeItem(i item.Item) {
	// TODO: Add max carry weight in the future
	p.Inventory = append(p.Inventory, i)
}
