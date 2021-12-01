package server

import "gopkg.in/ini.v1"

type ServerContext struct {
	Ini *ini.File
}

func NewServerContext() *ServerContext {
	return &ServerContext{}
}
