package models

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"src/pkg/utils"
	"strings"
	"time"
)

const (
	kTokenAge = 10 * time.Minute
)

// AuthJwtToken is a struct that holds the jwt token and the bearer token
type AuthJwtToken struct {
	BearerToken string `json:"bearer_token"`
	JwtToken    *jwt.Token
}

// NewAuthJwtToken creates a new jwt token and returns it
func NewAuthJwtToken(secret UserSecret) (AuthJwtToken, error) {
	var err error
	var tok AuthJwtToken
	tok.JwtToken = secret.CreateJWT(kTokenAge)
	tok.BearerToken, err = tok.JwtToken.SignedString(utils.GetAPISecretBytes())
	return tok, err
}

// ExtractAuthToken extracts the jwt token from the request header
func ExtractAuthToken(c *gin.Context) (string, error) {
	authHeader, found := c.Request.Header["Authorization"]
	if !found {
		return "", fmt.Errorf("authorization header not found")
	}
	// extract the jwt token
	authParts := strings.Split(authHeader[0], " ")
	if len(authParts) != 2 || authParts[0] != string(BearerAuth) {
		return "", fmt.Errorf("bad auth header")
	}
	return authParts[1], nil
}

// ParseJWTToken parses a token string and returns the jwt token
func ParseJWTToken(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, utils.JwtKeyFunc)
	return token, err
}
