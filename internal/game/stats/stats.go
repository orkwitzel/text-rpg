package stats

type Stats struct {
	Attack       int `json:"attack"`
	Defense      int `json:"defense"`
	Speed        int `json:"speed"`
	Intelligence int `json:"intelligence"`
	Luck         int `json:"luck"`
}

func NewStats(attack int, defense int, speed int, intelligence int, luck int) Stats {
	return Stats{Attack: CheckStatValue(attack),
		Defense:      CheckStatValue(defense),
		Speed:        CheckStatValue(speed),
		Intelligence: CheckStatValue(intelligence),
		Luck:         CheckStatValue(luck),
	}
}

func CheckStatValue(stat int) int {
	if stat < 0 {
		return 0
	} else if stat > 10 {
		return 10
	}
	return stat
}
