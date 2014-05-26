package main

import (
  "os"
  "koding/tools/logger"
)

var log = logger.New("name")

// By default log writes to stdout, file and syslog
// stdout/file have color, syslog doesn't

func main() {
  log.SetFileName("name")

  log.Critical("crit")
  log.Error("err")
  log.Warning("warning")
  log.Notice("notice")
  log.Info("info")
}
