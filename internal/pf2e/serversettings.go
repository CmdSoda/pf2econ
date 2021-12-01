package pf2e

import "gopkg.in/ini.v1"

type ServerSettings struct {
	DataPath string
}

func NewSettings(ini *ini.File) *ServerSettings {
	s := ServerSettings{}
	s.DataPath = ini.Section("").Key("dbpath").String()
	return &s
}
