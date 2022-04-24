package main

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

func main() {
	accessSecrte := "ae0536f9-6450-4606-8e13-5a19ed505da0"
	accessToken, err := buildToken(accessSecrte, map[string]interface{}{
		"key": "value",
	}, 60*5)

	if err != nil {
		fmt.Printf("gentoken err %v", err)
	}
	fmt.Println("accessToken:", accessToken)
}

func buildToken(secretKey string, payloads map[string]interface{}, seconds int64) (string, error) {
	now := time.Now().Unix()
	claims := make(jwt.MapClaims)
	claims["exp"] = now + seconds
	claims["iat"] = now
	for k, v := range payloads {
		claims[k] = v
	}
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	return token.SignedString([]byte(secretKey))
}
