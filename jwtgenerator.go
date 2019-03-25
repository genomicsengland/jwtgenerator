package main

import (
	"fmt"
	"github.com/dvsekhvalnov/jose2go"
	"os"
	"crypto/md5"
)

func main() {
	key := os.Args[1]
	secret := os.Args[2]
	oid := os.Args[3]
	fmt.Printf("%v", generateToken(key, secret, oid))
}

func generateToken(key string, secret string, oid string) string {
	data       := []byte(oid)
	jti        := fmt.Sprintf("%x", md5.Sum(data))
	payload    := `{"iss":"` + key + `", "jti":"` + jti + `", "oid":"` + oid + `"}`
	sharedKey  := []byte(secret)
	token, err := jose.Sign(payload, jose.HS256, sharedKey, jose.Header("typ", "JWT"))

	if err != nil {
		return err.Error()
	}

	return token
}