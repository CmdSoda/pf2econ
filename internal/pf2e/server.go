package pf2e

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
	"io/ioutil"
	"net"
	"os"
)

type Server struct {
	Ini *ini.File
	*ServerSettings
}

func NewServer(iniName string) (*Server, error) {
	s := Server{}
	err := s.init(iniName)
	return &s, err
}

func (s *Server) init(iniName string) error {
	if err := s.LoadIni(iniName); err != nil {
		return err
	}
	s.ServerSettings = NewSettings(s.Ini)
	return nil
}

func (s *Server) Run() {
	s.ListenAndAccept()
}

func (s *Server) LoadIni(iniName string) error {
	var err error
	s.Ini, err = ini.Load(iniName)
	if err != nil {
		log.Errorf("failed to read file: %v\n", err)
		return err
	}
	return nil
}

//goland:noinspection GoUnhandledErrorResult
func (s *Server) ListenAndAccept() {
	l, err := net.Listen("tcp", "localhost:"+s.Ini.Section("").Key("port").String())
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
func (s *Server) handleRequest(conn net.Conn) {
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

func (s *Server) LoadCharacter(filename string) (*Character, error) {
	c := Character{}
	dataPathFilename := s.DataPath + filename
	file, err := os.Open(dataPathFilename)
	if err != nil {
		log.Errorf("%s not found\n", dataPathFilename)
		return &c, err
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Errorf("%s error while reading\n", dataPathFilename)
		return &c, err
	}
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		log.Errorf("%s error while unmarshaling\n", dataPathFilename)
		return &c, err
	}
	return &c, nil
}

func (s *Server) SaveCharacter(character *Character, filename string) error {
	file, err := json.MarshalIndent(character, "", " ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(s.DataPath+filename, file, 0644)
	return err
}
