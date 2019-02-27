package main

import (
	"fmt"
	"github.com/dvsekhvalnov/jose2go"
	"os"
)

func main() {
	key := os.Args[1]
	secret := os.Args[2]
	fmt.Printf("%v", generateToken(key, secret))
}

func generateToken(key string, secret string) string {
	payload := `{"iss":"` + key + `"}`
	sharedKey := []byte(secret)
	token, err := jose.Sign(payload, jose.HS256, sharedKey, jose.Header("typ", "JWT"))

	if err != nil {
		return err.Error()
	}

	return token
}