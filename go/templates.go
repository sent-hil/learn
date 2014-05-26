package main

import (
	"html/template"
	"os"
)

func main() {
	f, err := os.Create("charts.html")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	t, err := template.ParseFiles("templates.html")
	if err != nil {
		panic(err)
	}

	var data = []interface{}{
		[]interface{}{23, 1388369100},
		[]interface{}{33, 1388369700},
		[]interface{}{30, 1388370000},
		[]interface{}{28, 1388370300},
		[]interface{}{28, 1388370600},
	}

	if err := t.Execute(f, map[string]interface{}{"data": data}); err != nil {
		panic(err)
	}
}
