package pf2e

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func newTestChar() *Character {
	c := NewCharacter()
	c.Name = "John"
	c.Name = "John"
	c.Abilities = AbilityScores{
		Strength:     1,
		Dexterity:    2,
		Constitution: 3,
		Intelligence: 4,
		Wisdom:       5,
		Charisma:     6,
	}
	return c
}

//goland:noinspection SpellCheckingInspection
var testChar = "testchar.json"
var testCharIni = "testing/test.ini"

func TestPf2Server_SaveCharacter(t *testing.T) {
	s, err := NewServer(testCharIni)
	assert.Nil(t, err)
	c := newTestChar()
	err = s.SaveCharacter(c, testChar)
	assert.Nil(t, err)
	fmt.Println(c)
}

func TestPf2Server_LoadCharacter(t *testing.T) {
	var err error
	var s *Server
	var c *Character
	s, err = NewServer(testCharIni)
	assert.Nil(t, err)
	c, err = s.LoadCharacter(testChar)
	assert.Nil(t, err)
	ts := newTestChar()
	assert.Equal(t, c.Name, ts.Name)
	assert.Equal(t, c.Abilities.Strength, ts.Abilities.Strength)
	assert.Equal(t, c.Abilities.Dexterity, ts.Abilities.Dexterity)
	assert.Equal(t, c.Abilities.Constitution, ts.Abilities.Constitution)
	assert.Equal(t, c.Abilities.Intelligence, ts.Abilities.Intelligence)
	assert.Equal(t, c.Abilities.Wisdom, ts.Abilities.Wisdom)
	assert.Equal(t, c.Abilities.Charisma, ts.Abilities.Charisma)
}
