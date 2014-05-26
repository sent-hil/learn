package main

import (
	"log"
	"runtime"
)

func main() {
	// will use all cores in machines
	// goroutines are not guranteed to run synchrously
	runtime.GOMAXPROCS(runtime.NumCPU())

	var _, file, line, ok = runtime.Caller(0)
	if ok {
		log.Println(file, line)
	}
}
