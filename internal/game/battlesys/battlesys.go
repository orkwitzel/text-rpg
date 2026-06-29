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
	if !playerRollsHit(p, e) {
		return 0
	}

	attackPower := (p.Stats.Attack * 4) + (p.Level * 3) + (p.Stats.Intelligence / 2) + weaponDamageBonus(p.EquippedWeapon)
	mitigation := (e.Stats.Defense * 3) + (e.Level * 2)
	damage := attackPower - mitigation/2
	if damage < 1 {
		damage = 1
	}

	damage = applyVariance(damage)
	if playerRollsCrit(p) {
		damage = damage * 3 / 2
	}

	return damage
}

func CalculateDamageToPlayer(p *player.Player, e *enemy.Enemy) int {
	if !enemyRollsHit(e, p) {
		return 0
	}

	attackPower := (e.Stats.Attack * 4) + (e.Level * 3) + (e.Stats.Intelligence / 2)
	mitigation := (p.Stats.Defense * 3) + (p.Level * 2)
	damage := attackPower - mitigation/2
	if damage < 1 {
		damage = 1
	}

	damage = applyVariance(damage)
	if enemyRollsCrit(e) {
		damage = damage * 3 / 2
	}

	return damage
}

func enemyRollsCrit(e *enemy.Enemy) bool {
	critChance := clamp(e.Stats.Luck*5, 0, maxCritRate)
	return rand.Intn(100) < critChance
}

func enemyRollsHit(e *enemy.Enemy, p *player.Player) bool {
	accuracy := e.Stats.Luck*3 + e.Stats.Speed*2 + e.Level*2
	evasion := p.Stats.Speed*3 + p.Stats.Luck*2 + p.Level*2
	hitRate := clamp(50+accuracy-evasion, minHitRate, maxHitRate)
	return rand.Intn(100) < hitRate
}

func playerRollsHit(p *player.Player, e *enemy.Enemy) bool {
	accuracy := p.Stats.Luck*3 + p.Stats.Speed*2 + p.Level*2
	evasion := e.Stats.Speed*3 + e.Stats.Luck*2 + e.Level
	hitRate := clamp(50+accuracy-evasion, minHitRate, maxHitRate)
	return rand.Intn(100) < hitRate
}

func playerRollsCrit(p *player.Player) bool {
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
