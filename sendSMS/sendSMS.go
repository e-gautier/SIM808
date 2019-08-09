package main

import (
	"bytes"
	"encoding/hex"
	"log"
	"log/syslog"
	"os"
	"strconv"
	"strings"

	"github.com/xlab/at/sms"

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

	message := strings.Join(args[2:], " ")

	// set SMS message format to PDU (0) mode
	serialwrapper.Send(s, "AT+CMGF=0")

	smsPDU := new(sms.Message)
	smsPDU.Text = message
	smsPDU.Address = sms.PhoneNumber(number)
	smsPDU.Type = sms.MessageTypes.Submit
	length, data, err := smsPDU.PDU()
	if err != nil {
		log.Fatal(err)
	}

	hexaPDU := hex.EncodeToString(data)
	var buffer bytes.Buffer
	buffer.WriteString("AT+CMGS=")
	buffer.WriteString(strconv.Itoa(length))

	serialwrapper.Send(s, buffer.String())
	serialwrapper.Send(s, hexaPDU)
	serialwrapper.Send(s, "\u001A")

	// syslog writer
	buffer.Reset()
	buffer.WriteString("sms sent to ")
	buffer.WriteString(number)
	buffer.WriteString(" ")
	buffer.WriteString(message)
	_ = syslogWriter.Info(buffer.String())
}
