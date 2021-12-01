package main

import (
	"github.com/CmdSoda/pf2econ/internal/pf2e"
	log "github.com/sirupsen/logrus"
)

func main() {
	c, err := pf2e.NewClient("pf2econ.ini")
	if err != nil {
		log.Errorln("failed to create client")
		return
	}
	c.Run()
}
