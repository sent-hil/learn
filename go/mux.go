package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var Current Status

type Worker struct {
	name string
}

type Status struct {
	RegisteredWorkers map[string]Worker
}

type MatcherFunc func(*http.Request, *mux.RouteMatch) bool

func (m MatcherFunc) Match(r *http.Request, match *mux.RouteMatch) bool {
	return m(r, match)
}

func main() {
	one := Worker{"one"}
	two := Worker{"two"}
	boo := true
	workers := map[string]Worker{"1": one, "2": two}
	Current = Status{workers}

	log.Println(HelloFunc(func(a int) {}).Hi())

	b := MatcherFunc(func(r *http.Request, m *mux.RouteMatch) bool { return boo })
	log.Println(b.Match())

	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)

	http.Handle("/", r)
	log.Println("Starting server on :8001")

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Println("Couldn't start server", err)
	}
}

func homeHandler(resp http.ResponseWriter, req *http.Request) {
	log.Println("Got request on /\n")
}

type HelloFunc func(a int)

func (h HelloFunc) Hi() bool {
	return true
}

//func customMatcher() func(*http.Request, *mux.RouteMatch) bool {
//return func(*http.Request, *mux.RouteMatch) bool {
//return true
//}
//}
