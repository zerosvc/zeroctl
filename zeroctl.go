package main

import (
	//	"fmt"
	"github.com/op/go-logging"
	"os"
)

var log = logging.MustGetLogger("example")

//var format = "%{color}%{time:2006-01-02T15:04:05.9999Z-07:00} → %{level:.4s} %{id:03x}%{color:reset} %{message}"
var format = "%{color:bold}%{time:2006-01-02T15:04:05.9999Z-07:00}%{color:reset}%{color} [%{level:.1s}] %{color:reset}%{shortpkg}[%{longfunc}] %{message}"

func main() {
	logBackend := logging.NewLogBackend(os.Stderr, "", 0)
	syslogBackend, err := logging.NewSyslogBackend("")
	if err != nil {
		log.Fatal(err)
	}
	logging.SetBackend(logBackend, syslogBackend)
	logging.SetFormatter(logging.MustStringFormatter(format))
	log.Error("Logging")
	log.Warning("Works")

}