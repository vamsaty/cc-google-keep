package utils

import (
	"crypto/sha256"
	"fmt"
	"github.com/golang-jwt/jwt"
	"os"
)

var (
	dummySecret = "SecretYouShouldHide"

	JwtKeyFunc = func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return GetAPISecretBytes(), nil
	}
)

func GetAPISecretBytes() []byte {
	secret, found := os.LookupEnv("API_SECRET")
	if !found {
		return []byte(dummySecret)
	}
	return []byte(secret)
}

func SimpleHash(pwd string) string {
	h := sha256.New()
	h.Write([]byte(pwd))
	return string(h.Sum(nil))
}
