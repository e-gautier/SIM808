package main

import (
	"log"
	"time"

	"github.com/tarm/serial"
)

func main() {

	c := &serial.Config{Name: "/dev/ttyAMA0", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	for {
		buf := make([]byte, 128)
		n, err := s.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(string(buf))
		log.Printf("%s", string(buf[:n]))

		time.Sleep(2 * time.Second)
	}

}
