package pf2e

import (
	"fmt"
	"github.com/jroimartin/gocui"
	log "github.com/sirupsen/logrus"
)

type Client struct {
}

func NewClient(iniName string) (*Client, error) {
	c := Client{}

	return &c, nil
}

func (c *Client) Run() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

var reset = "\u001b[0m"
var white = "\u001b[37;1m"
var yellow = "\u001b[33;1m"

func layout(g *gocui.Gui) error {
	//maxX, maxY := g.Size()
	if v, err := g.SetView("hello", 0, 0, 40, 20); err != nil {
		v.Title = "Ability Scores"
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, white+"Str: "+yellow+"12")
		fmt.Fprintln(v, white+"Dex: "+yellow+"13")
		fmt.Fprintln(v, white+"Con: "+yellow+"11")
		fmt.Fprintln(v, white+"Int: "+yellow+"14")
		fmt.Fprintln(v, white+"Wis: "+yellow+"15")
		fmt.Fprintln(v, white+"Cha: "+yellow+"14")
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
