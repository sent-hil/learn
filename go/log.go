package main

import (
	"fmt"
	"github.com/op/go-logging"
	stdlog "log"
	"os"
)

func New(name string) *logging.Logger {
	logging.SetFormatter(logging.MustStringFormatter("[%{level:.8s}] - %{message}"))

	var logBackend = logging.NewLogBackend(os.Stderr, "", stdlog.LstdFlags|stdlog.Lshortfile)
	logBackend.Color = true

	var syslogBackend, err = logging.NewSyslogBackend("")
	if err != nil {
		panic(err)
	}

	logging.SetBackend(logBackend, syslogBackend)

	return logging.MustGetLogger(name)
}

func main() {
	var log = New("hola")

	log.Critical("crit")
	log.Error("err")
	log.Warning("warning")
	log.Notice("notice")
	log.Info("info")

	var err error

	fmt.Println(fmt.Sprintf("Error in Read of an int16: %s (%d bytes read)\n", err, 10), 0)

	defer func() {
		if err := recover(); err != nil {
			log.Critical("Panicked %v", err)
		}
	}()

	panic("hello")
}
