package main

import (
	"github.com/CmdSoda/pf2econ/internal/server"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Infoln("Pathfinder 2e Console Server")
	s, err := server.NewServer("pf2econs.ini")
	if err != nil {
		log.Panicln(err.Error())
	}
	s.Run()
}
