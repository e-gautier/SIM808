package serialwrapper

import (
	"bytes"
	"log"

	"github.com/tarm/serial"
)

// Send comand through serial port
func Send(s *serial.Port, command string) {
	var buffer bytes.Buffer
	buffer.WriteString(command)
	buffer.WriteString("\n")
	b := []byte(buffer.String())
	_, err := s.Write(b)
	if err != nil {
		log.Fatal(err)
	}
}

// Init serial port
func Init() *serial.Port {
	c := &serial.Config{Name: "/dev/ttyAMA0", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	return s
}
