package main

import (
	"fmt"
	"os"
)

func main() {
	process, _ := os.FindProcess(6255)
	fmt.Println(process)

	process.Kill()
}
