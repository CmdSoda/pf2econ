package server

import (
	"encoding/json"
	"fmt"
	"github.com/CmdSoda/pf2econ/internal/game"
	log "github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
	"io/ioutil"
	"net"
	"os"
)

type Pf2Server struct {
	*ServerContext
	*Settings
}

func NewServer(iniName string) (*Pf2Server, error) {
	s := Pf2Server{}
	err := s.init(iniName)
	return &s, err
}

func (s *Pf2Server) init(ininame string) error {
	s.ServerContext = NewServerContext()
	if err := s.LoadIni(ininame); err != nil {
		return err
	}
	s.Settings = NewSettings(s.Ini)
	return nil
}

func (s *Pf2Server) Run() {
	s.ListenAndAccept()
}

func (s *Pf2Server) LoadIni(iniName string) error {
	var err error
	s.Ini, err = ini.Load(iniName)
	if err != nil {
		log.Errorf("failed to read file: %v\n", err)
		return err
	}
	return nil
}

//goland:noinspection GoUnhandledErrorResult
func (s *Pf2Server) ListenAndAccept() {
	l, err := net.Listen("tcp", "localhost:"+s.ServerContext.Ini.Section("").Key("port").String())
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go s.handleRequest(conn)
	}
}

//goland:noinspection GoUnhandledErrorResult
func (s *Pf2Server) handleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	// Send a response back to person contacting us.
	conn.Write([]byte("Message received.\n"))
	// Close the connection when you're done with it.
	conn.Close()
}

func (s *Pf2Server) LoadCharacter(filename string) (game.Character, error) {
	c := game.Character{}
	dataPathFilename := s.DataPath + filename
	file, err := os.Open(dataPathFilename)
	if err != nil {
		log.Errorf("%s not found\n", dataPathFilename)
		return c, err
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Errorf("%s error while reading\n", dataPathFilename)
		return c, err
	}
	err = json.Unmarshal(bytes, &s)
	if err != nil {
		log.Errorf("%s error while unmarshaling\n", dataPathFilename)
		return c, err
	}
	return c, nil
}

func (s *Pf2Server) SaveCharacter(character *game.Character, filename string) error {
	file, err := json.MarshalIndent(character, "", " ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(s.DataPath+filename, file, 0644)
	return err
}
