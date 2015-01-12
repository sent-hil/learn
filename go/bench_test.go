package main

import (
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"log"
	"testing"
	"time"
)

func BenchmarkHs(*testing.B) {
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims["foo"] = "bar"
	token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	token.SignedString([]byte{0, 0, 0, 0, 0})
}

func BenchmarkRs(*testing.B) {
	var pp []byte

	privateKey, err := ioutil.ReadFile("/tmp/privateKey.pem")
	if err != nil {
		log.Fatalf("cannot read public key file: %s", err.Error())
	}

	pp = privateKey

	token := jwt.New(jwt.GetSigningMethod("RS256"))
	token.Claims["iss"] = "issuer"   // Issuer
	token.Claims["sub"] = "user"     // Subject
	token.Claims["aud"] = "audience" // Audience
	token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	token.SignedString(pp)
}
