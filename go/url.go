package main

import "fmt"

func main() {
	var name = "summarize(keepLastValue(perfdata.kontainer.kontainer1.sj.koding.com.vm.online), '10min', 'avg')"

	fmt.Println(fmt.Sprintf("http://graphite.sj.koding.com/render/?target=%s&format=json", name))
}
