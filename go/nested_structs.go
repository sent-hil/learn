package main

import (
	"log"
)

type TokenData struct {
	name string
}

type Token interface {
	GetServiceName() string
}

type TwitterToken struct {
	*TokenData
}

type GithubToken struct {
	*TokenData
}

func (t *TwitterToken) GetServiceName() string {
	return t.name + "twitter"
}

func (t *GithubToken) GetServiceName() string {
	return t.name + "github"
}

func namesToConstructors(tokenName string) Token {
	vals := map[string](func(string) Token){
		"github":  NewGithub,
		"twitter": NewTwitter,
	}

	token := vals[tokenName](tokenName)

	return token
}

func main() {
	twitter := namesToConstructors("twitter")
	log.Println(twitter.GetServiceName())

	github := namesToConstructors("github")
	log.Println(github.GetServiceName())
}

func NewGithub(name string) Token {
	data := &TokenData{"from "}
	github := &GithubToken{data}

	return github
}

func NewTwitter(name string) Token {
	data := &TokenData{"from "}
	twitter := &TwitterToken{data}

	return twitter
}
