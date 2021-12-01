package server

import (
	"fmt"
	"github.com/CmdSoda/pf2econ/internal/game"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPf2Server_SaveCharacter(t *testing.T) {
	s := Pf2Server{}
	assert.Nil(t, s.Init("testing/test.ini"))
	c := game.NewCharacter()
	err := s.SaveCharacter(c, "testchar2.json")
	assert.Nil(t, err)
	fmt.Println(c)
}

func TestPf2Server_LoadCharacter(t *testing.T) {
	s := Pf2Server{}
	assert.Nil(t, s.Init("testing/test.ini"))
	c, err := s.LoadCharacter("testchar2.json")
	assert.Nil(t, err)
	fmt.Println(c)
}
