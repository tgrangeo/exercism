package dndcharacter

import (
	"math"
	"math/rand"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

func Modifier(score int) int {
	return int(math.Floor(float64(score-10) / 2.0))
}

func Ability() int {
	return rand.Intn(15) + 3
}

func GenerateCharacter() Character {
	str := Ability()
	dex := Ability()
	con := Ability()
	intel := Ability()
	wis := Ability()
	cha := Ability()
	hit := 10 + Modifier(con)
	return Character{str, dex, con, intel, wis, cha, hit}
}
