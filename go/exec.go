package main

import (
	"bytes"
	"log"
	"os/exec"
)

func main() {
	box := "1"
	version := "2"
	cmd := exec.Command("echo", box, version)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%q", out.String())
}
