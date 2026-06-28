package battlesys

import (
	"math/rand"
	"rpg/internal/game/item"
	"rpg/internal/game/npcs/enemy"
	"rpg/internal/game/player"
)

const (
	minHitRate  = 5
	maxHitRate  = 95
	maxCritRate = 25
)

func CalculateDamageToEnemy(e *enemy.Enemy, p *player.Player) int {
	if !rollHit(p, e) {
		return 0
	}

	attackPower := (p.Stats.Attack * 4) + (p.Level * 3) + (p.Stats.Intelligence / 2) + weaponDamageBonus(p.EquippedWeapon)
	mitigation := (e.Stats.Defense * 3) + (e.Level * 2)
	damage := attackPower - mitigation/2
	if damage < 1 {
		damage = 1
	}

	damage = applyVariance(damage)
	if rollCrit(p) {
		damage = damage * 3 / 2
	}

	return damage
}

func rollHit(p *player.Player, e *enemy.Enemy) bool {
	accuracy := p.Stats.Luck*3 + p.Stats.Speed*2 + p.Level*2
	evasion := e.Stats.Speed*3 + e.Stats.Luck*2 + e.Level
	hitRate := clamp(50+accuracy-evasion, minHitRate, maxHitRate)
	return rand.Intn(100) < hitRate
}

func rollCrit(p *player.Player) bool {
	critChance := clamp(p.Stats.Luck*5, 0, maxCritRate)
	return rand.Intn(100) < critChance
}

func weaponDamageBonus(weapon *item.Item) int {
	if weapon == nil {
		return 0
	}
	bonus := 0
	for _, effect := range weapon.Effects {
		bonus += effect.Damage
	}
	return bonus
}

func applyVariance(damage int) int {
	multiplier := 85 + rand.Intn(31)
	return damage * multiplier / 100
}

func clamp(v, min, max int) int {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}
