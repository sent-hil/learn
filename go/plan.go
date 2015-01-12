package main

import (
	"fmt"
	"io"
	"net/http"
)

var plan = "free"

func main() {
	http.HandleFunc("/", checkerHttp)
	http.ListenAndServe(":"+"7000", nil)
}

func checkerHttp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	accountId := r.URL.Query().Get("account_id")
	if accountId == "" {
		io.WriteString(w, `{"error":"account_id is required"}`)
		return
	}

	fmt.Printf(">>>>>>>> Returning %s\n", plan)

	io.WriteString(w, fmt.Sprintf(`{"planTitle":"%s"}`, plan))
}
