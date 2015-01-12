package main

import (
	"os"
	"os/signal"
	"runtime/pprof"
)

func fib(num int) (result int) {
	if num > 2 {
		result = fib(num-1) + fib(num-2)
	} else {
		result = 1
	}

	return
}

func main() {
	f, _ := os.Create("fib.prof")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for _ = range c {
			pprof.StopCPUProfile()
			os.Exit(0)
		}
	}()

	for i := 1; i < 40; i++ {
		fib(i)
	}
}
