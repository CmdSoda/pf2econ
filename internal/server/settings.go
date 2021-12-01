package server

import "gopkg.in/ini.v1"

type Settings struct {
	DataPath string
}

func NewSettings(ini *ini.File) *Settings {
	s := Settings{}
	s.DataPath = ini.Section("").Key("datapath").String()
	return &s
}
