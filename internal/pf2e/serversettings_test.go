package pf2e

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/ini.v1"
	"testing"
)

func TestNewSettings(t *testing.T) {
	var err error
	var i *ini.File
	i, err = ini.Load("testing/test.ini")
	assert.Nil(t, err)
	s := NewSettings(i)
	assert.Equal(t, s.DataPath, "../../db/")
}
