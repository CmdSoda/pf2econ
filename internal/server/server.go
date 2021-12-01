package server

import (
	"fmt"
	"gopkg.in/ini.v1"
	"net"
	"os"
)

type Pf2Server struct {
	*ServerContext
}

func (s *Pf2Server) LoadIni() {
	var err error
	s.Configuration, err = ini.Load("pf2econs.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
}

//goland:noinspection GoUnhandledErrorResult
func (s *Pf2Server) ListenAndAccept() {
	l, err := net.Listen("tcp", "localhost:"+s.ServerContext.Configuration.Section("").Key("port").String())
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
