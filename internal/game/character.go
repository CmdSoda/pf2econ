package game

type AbilityScores struct {
	Strength     uint32
	Dexterity    uint32
	Constitution uint32
	Intelligence uint32
	Wisdom       uint32
	Charisma     uint32
}

type Charakter struct {
	Name string
	AbilityScores
}
