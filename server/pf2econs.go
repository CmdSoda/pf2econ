package main

import (
	"github.com/CmdSoda/pf2econ/internal/server"
	log "github.com/sirupsen/logrus"
)

func main() {
	s := server.Pf2Server{}
	log.Infoln("Pathfinder 2e Console Server")
	s.ServerContext = server.NewServerContext()
	s.LoadIni()
	s.ListenAndAccept()
}
