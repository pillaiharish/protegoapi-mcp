package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	secret := []byte("I_AM_GROOT")
	claims := jwt.MapClaims{
		"sub":   "harish",
		"roles": []string{"admin"},
		"exp":   time.Now().Add(time.Hour).Unix(),
	}
	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := tok.SignedString(secret)
	if err != nil {
		panic(err)
	}
	fmt.Println(signed)
}

