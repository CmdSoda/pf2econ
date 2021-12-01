package main

import "gopkg.in/ini.v1"

type ServerContext struct {
	Configuration *ini.File
}

func NewServerContext() *ServerContext {
	return &ServerContext{}
}
