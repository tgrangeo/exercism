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

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	return 0
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	return 0
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	c := Character{
		Strength:     rand.Intn(15) + 3,
		Dexterity:    rand.Intn(15) + 3,
		Constitution: rand.Intn(15) + 3,
		Intelligence: rand.Intn(15) + 3,
		Wisdom:       rand.Intn(15) + 3,
		Charisma:     rand.Intn(15) + 3,
		Hitpoints:    0,
	}
	c.Hitpoints = int(math.Round(float64(c.Constitution - 10) / 2))
	return c
}
