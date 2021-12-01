package pf2e

import "github.com/jroimartin/gocui"

type CharacterView struct {
	Name   string
	Title  string
	Left   int
	Top    int
	Width  int
	Height int
	View   *gocui.View
}

// CharacterPage is a collection of CharacterView objects.
type CharacterPage struct {
	Views []*CharacterView
}

// NewCharacterPage creates a Page from a definition file.
//goland:noinspection GoUnusedParameter,GoUnusedExportedFunction
func NewCharacterPage(name string) *CharacterPage {
	return &CharacterPage{}
}
