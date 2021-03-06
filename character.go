package main

import (
	"fmt"
	"math/rand"
)

type DiceRoll struct {
	Sides    int `json:"sides"`
	NumRolls int `json:"rolls"`
}

func (d *DiceRoll) Roll() int {
	sum := 0
	for i := 0; i < d.NumRolls; i += 1 {
		roll := (rand.Intn(d.Sides) + 1)
		fmt.Println("Roll " + string(i) + " is " + string(roll))
		sum += roll
	}

	return sum
}

type Character struct {
	Name           string      `json:"name"`
	Race           Race        `json:"race"`     //Making it a pointer is more effiecient because there are a
	Class          Class       `json:"class"`    //set number of instatiated races and classes that we can just reference
	Portrait       string      `json:"portrait"` //Maybe an image name???
	Description    Description `json:"desc"`
	Exp            int         `json:"exp"`
	Stats          Stats       `json:"stats"`
	Modifiers      Modifiers   `json:"mods"`
	Inventory      []Item      `json:"inventory"`
	EquipedClothes []Wearable  `json:"clothes"`
	EquipedWeapon  []Weapon    `json:"weapons"`
	Skills         []Skill     `json:"skills"`
	Spells         []Spell     `json:"spells"`
}

type Race struct {
	Name        string    `json:"name"`
	Image       string    `json:"image"`
	Description string    `json:"desc"` //Possible front end use???
	Traits      []string  `json:"traits"`
	Bonuses     Modifiers `json:"bonuses"`
}

type Class struct {
	Name          string   `json:"name"`
	Image         string   `json:"image"`
	Subclass      string   `json:"subclass"`
	Description   string   `json:"desc"` //Possible front end use???
	Proficiencies []string `json:"proficiencies"`
}

type Description struct {
	Alignment  string `json:"alignment"`
	Ideals     string `json:"ideals"`
	Bonds      string `json:"bonds"`
	Flaws      string `json:"flaws"`
	Background string `json:"background"`
}

type Stats struct {
	Strength     int `json:"strength"`
	Dexterity    int `json:"dexterity"`
	Constitution int `json:"constitution"`
	Intelligence int `json:"intelligence"`
	Wisdom       int `json:"wisdom"`
	Charisma     int `json:"charisma"`
}

type Modifiers struct {
	StrengthMod     int `json:"strengthMod"`
	DexterityMod    int `json:"dexterityMod"`
	ConstitutionMod int `json:"constitutionMod"`
	IntelligenceMod int `json:"intelligenceMod"`
	WisdomMod       int `json:"wisdomMod"`
	CharismaMod     int `json:"charismaMod"`
}

type Item struct {
	Name        string `json:"name"`
	Worth       int    `json:"worth"` //Worth in GP
	Description string `json:"desc"`
}

type Wearable struct {
	Item
	AC int `json:"ac"`
}

type Weapon struct {
	Item
	WeaponType string   `json:"type"`
	Damage     DiceRoll `json:"damage"` //Possibly create a struct of some kind to represent dicerolls???
}

type Skill struct {
	Name        string `json:"name"`
	PrimaryStat string `json:"primaryStat"`
}

type Spell struct {
	Name        string `json:"name"`
	Description string `json:"desc"`
	Type        string `json:"type"` //Defines what type of skill i.e. usable per day/hour/etc
	MaxUses     int    `json:"numUses"`
}

func NewMod(myStats Stats) Modifiers {
	mod := Modifiers{
		StrengthMod:     (myStats.Strength - 10) / 2,
		DexterityMod:    (myStats.Dexterity - 10) / 2,
		ConstitutionMod: (myStats.Constitution - 10) / 2,
		IntelligenceMod: (myStats.Intelligence - 10) / 2,
		WisdomMod:       (myStats.Wisdom - 10) / 2,
		CharismaMod:     (myStats.Charisma - 10) / 2,
	}

	return mod
}
