package main

import (
	"bytes"
	"log"
	"os"
	"strings"

	"github.com/tarm/serial"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		log.Println("parameters not valid, require number, text")
		log.Println("sendSMS NUMBER TEXT")
		log.Println("ex: sendSMS 0123456789 hello World!")
		return
	}

	number := args[1]

	c := &serial.Config{Name: "/dev/ttyAMA0", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	send(s, "AT+CMGF=1")

	var numberBuffer bytes.Buffer
	numberBuffer.WriteString("AT+CMGS=\"")
	numberBuffer.WriteString(number)
	numberBuffer.WriteString("\"")
	send(s, numberBuffer.String())
	send(s, strings.Join(args[2:], " "))
	send(s, "\u001A")

	buf := make([]byte, 128)
	n, err := s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%q", buf[:n])

}

func send(s *serial.Port, command string) {
	var buffer bytes.Buffer
	buffer.WriteString(command)
	buffer.WriteString("\n")
	b := []byte(buffer.String())
	_, err := s.Write(b)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%q", b)
}
