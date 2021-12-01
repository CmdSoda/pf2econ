package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	s := Pf2Server{}
	log.Infoln("Pathfinder 2e Console Server")
	s.ServerContext = NewServerContext()
	s.LoadIni()
	s.ListenAndAccept()
}
