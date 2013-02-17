package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd, _ := exec.Command("/bin/sh", "-c", "ps aux | grep social").Output()
	fmt.Println(string(cmd))
}
