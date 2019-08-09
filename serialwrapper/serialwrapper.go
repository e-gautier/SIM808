package serialwrapper

import (
	"bytes"
	"log"
	"time"

	"github.com/tarm/serial"
)

const (
	serialDeviceFile = "/dev/ttyAMA0"
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
	time.Sleep(100 * time.Millisecond)
}

// Read serial output
func Read(s *serial.Port) string {
	buf := make([]byte, 128)
	n, err := s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	return string(buf[:n])
}

// Init serial port
func Init() *serial.Port {
	c := &serial.Config{Name: serialDeviceFile, Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	return s
}
