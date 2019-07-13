package config

import "fmt"

// SysLogTag global syslog tag
const SysLogTag string = "SMS808"

// SysLogTagReceiver sub tag for receiver
var SysLogTagReceiver = fmt.Sprintf("[%s] RECEIVER", SysLogTag)

// SysLogTagEmitter sub tag for emitter
var SysLogTagEmitter = fmt.Sprintf("[%s] EMITTER", SysLogTag)
