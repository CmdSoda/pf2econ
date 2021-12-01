package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
	"os"
)

var cfg *ini.File

func main() {
	log.Infoln("Pathfinder 2e Console")
	loadConfig()
	fmt.Printf("server = %s\n", cfg.Section("").Key("server").String())
	fmt.Printf("port = %s\n", cfg.Section("").Key("port").String())
}

func loadConfig() {
	var err error
	cfg, err = ini.Load("pf2econ.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
}
