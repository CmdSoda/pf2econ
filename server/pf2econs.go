package main

import (
	"github.com/CmdSoda/pf2econ/internal/server"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Infoln("Pathfinder 2e Console Server")
	s := server.Pf2Server{}
	s.Run()
}
