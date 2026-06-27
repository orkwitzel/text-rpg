package item

import "github.com/google/uuid"

type ItemType string

const (
	ItemTypeWeapon     ItemType = "weapon"
	ItemTypeArmor      ItemType = "armor"
	ItemTypeConsumable ItemType = "consumable"
	ItemTypeMaterial   ItemType = "material"
)

type Effect struct {
	Health int `json:"health"`
	Energy int `json:"energy"`
	Damage int `json:"damage"`
	Speed  int `json:"speed"`
	Level  int `json:"level"`
}

func NewEffect(health int, energy int, damage int, speed int, level int) *Effect {
	return &Effect{Health: health, Energy: energy, Damage: damage, Speed: speed, Level: level}
}

type Item struct {
	ID              string   `json:"id"`
	Name            string   `json:"name"`
	Description     string   `json:"description"`
	Value           int      `json:"value"`
	Weight          int      `json:"weight"`
	CanHaveMultiple bool     `json:"can_have_multiple"`
	Quantity        int      `json:"quantity"`
	ItemType        ItemType `json:"item_type"`
	Effects         []Effect `json:"effects"`
}

func NewItem(name string, description string, value int, weight int, canHaveMultiple bool, quantity int) *Item {
	return &Item{ID: uuid.New().String(), Name: name, Description: description, Value: value, Weight: weight, CanHaveMultiple: canHaveMultiple, Quantity: quantity}
}

func (i *Item) AddQuantity(quantity int) {
	if i.CanHaveMultiple {
		i.Quantity += quantity
	} else {
		i.Quantity = 1
	}
}

func (i *Item) RemoveQuantity(quantity int) {
	i.Quantity -= quantity
	if i.Quantity < 0 {
		i.Quantity = 0
	}
}
