package main

import "log"

type User struct {
	Name string
}

func main() {
	u := &User{Name: "Leto"}
	log.Printf("%p", u)

	log.Println(u.Name)
	Modify(&u)

	log.Println(u.Name)
}

func Modify(u **User) {
	*u = &User{Name: "Paul"}

	log.Printf("Modify: %p", u)
}
