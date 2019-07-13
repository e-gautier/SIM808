package main

import (
	"bytes"
	"log"
	"log/syslog"
	"os"
	"strings"

	"../config"
	"../serialwrapper"
)

// global syslog writer
var syslogWriter, _ = syslog.New(syslog.LOG_USER, config.SysLogTagEmitter)

func main() {
	args := os.Args
	if len(args) < 3 {
		log.Println("parameters not valid, require number, text")
		log.Println("sendSMS NUMBER TEXT")
		log.Println("ex: sendSMS 0123456789 hello World!")
		return
	}

	number := args[1]

	// init serial port
	s := serialwrapper.Init()
	_ = syslogWriter.Info("serial port initialized")

	// set SMS message format to text (1) mode
	serialwrapper.Send(s, "AT+CMGF=1")

	message := strings.Join(args[2:], " ")

	// send SMS message
	var buffer bytes.Buffer
	buffer.WriteString("AT+CMGS=\"")
	buffer.WriteString(number)
	buffer.WriteString("\"")
	serialwrapper.Send(s, buffer.String())
	serialwrapper.Send(s, message)
	serialwrapper.Send(s, "\u001A")

	// syslog writer
	buffer.Reset()
	buffer.WriteString("sms sent to ")
	buffer.WriteString(number)
	buffer.WriteString(" ")
	buffer.WriteString(message)
	_ = syslogWriter.Info(buffer.String())
}
