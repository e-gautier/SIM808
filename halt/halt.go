package main

import (
	"log"
	"syscall"

	rpio "github.com/stianeikeland/go-rpio"
)

const (
	pin = 13
)

func main() {
	err := rpio.Open()
	if err != nil {
		log.Fatal(err)
		syscall.Exit(-1)
	}
	pin := rpio.Pin(pin)

	pin.Output()
	pin.High()
	pin.Low()
}
