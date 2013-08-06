package main

import (
	"log"
	"runtime"
)

func main() {
	// will use all cores in machines
	// goroutines are not guranteed to run synchrously
	runtime.GOMAXPROCS(runtime.NumCPU())

	log.Println(10)
}
