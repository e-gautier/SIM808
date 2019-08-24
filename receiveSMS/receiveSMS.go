package main

import (
	"bytes"
	"encoding/hex"
	"log"
	"log/syslog"
	"os"
	"os/signal"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	"time"

	"../config"
	"../serialwrapper"
	"github.com/tarm/serial"
	"github.com/xlab/at/sms"
)

// for some reasons \n doesn't match in go so we need to use (?s) flag and replace \n with two random chars: .{2}
// original regex:
// \+CMGL: \d+,\d+,"\w*",\d+\n([[:xdigit:]]+)\n
var smsRegex = regexp.MustCompile(`(?s)\+CMGL: (\d+),(\d+),"\w*",\d+.{2}([[:xdigit:]]+)`)

// global buffer to handle the serial output
var bufferStr bytes.Buffer

// global syslog writer
var syslogWriter, _ = syslog.New(syslog.LOG_USER, config.SysLogTagReceiver)

func main() {
	// init the serial port
	s := serialwrapper.Init()
	defer s.Close()
	_ = syslogWriter.Info("serial port initialized")

	serialwrapper.Send(s, "AT+CMGF=0")

	// start to continually read the serial output and fill the buffer with it
	go read(s)

	go handleMessages(s)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	signal := <-c
	log.Println(signal)
	syscall.Exit(0)
}

// read the serial output and put it to the global buffer
func read(s *serial.Port) {
	for {
		output := serialwrapper.Read(s)
		bufferStr.WriteString(output)
	}
}

func handleMessages(s *serial.Port) {
	// indefinitely
	for {
		// tell to the chip that we want all messages with the CMGL command
		requestMessages(s)

		// sort by regex the buffer content to get the messages indexes and content
		smsmeta := smsRegex.FindAllStringSubmatch(bufferStr.String(), -1)

		// for each message
		for i := 0; i < len(smsmeta); i++ {
			// decode the message content
			rawsms, err := hex.DecodeString(smsmeta[i][3])
			if err != nil {
				continue
			}

			// attempt to create a message struct from the byte array
			message := new(sms.Message)
			_, err = message.ReadFrom(rawsms)
			if err != nil {
				continue
			}

			// build the log
			var log strings.Builder
			log.WriteString(string(message.Address))
			log.WriteString(", ")
			log.WriteString(message.Text)

			// write in log the message
			_ = syslogWriter.Info(log.String())

			// get the index and request a deletion of it
			index, _ := strconv.Atoi(smsmeta[i][1])
			delete(s, index)
		}

		// clear the buffer
		bufferStr.Reset()
	}
}

// request all messages from the chip with the CMGL command
func requestMessages(s *serial.Port) {
	var numberBuffer bytes.Buffer
	numberBuffer.WriteString("AT+CMGL=4")
	serialwrapper.Send(s, numberBuffer.String())
	serialwrapper.Send(s, "\u001A")
	time.Sleep(1 * time.Second)
}

// send a delete command according to an index
func delete(s *serial.Port, index int) {
	var numberBuffer bytes.Buffer
	numberBuffer.WriteString("AT+CMGD=")
	numberBuffer.WriteString(strconv.Itoa(index))
	serialwrapper.Send(s, numberBuffer.String())
	serialwrapper.Send(s, "\u001A")
	time.Sleep(1 * time.Second)
}
