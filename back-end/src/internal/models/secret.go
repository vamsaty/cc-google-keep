package models

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type AuthType string

const (
	BearerAuth AuthType = "Bearer"
)

// UserSecret is a struct that holds the user's credentials
type UserSecret struct {
	// ID and Username use the same values
	ID       string `json:"_id" bson:"_id"`
	Username string `json:"username" bson:"username"`
	// Password is received in plain text from the client and is hashed before
	// being stored in the database
	Password string `json:"password" bson:"password"`
}

// CreateJWT creates a jwt token with the user's id and username
func (secret *UserSecret) CreateJWT(expTime time.Duration) *jwt.Token {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"exp":        time.Now().Add(expTime).Unix(),
			"authorized": true,
			"user_id":    secret.ID, // as stored in db
			"username":   secret.Username,
		},
	)
	return token
}

func NewUserSecret(id, password string) UserSecret {
	return UserSecret{
		ID:       id,
		Username: id,
		Password: password,
	}
}

func NewSecretFromUser(user *User) UserSecret {
	return UserSecret{
		ID:       user.Username,
		Username: user.Username,
	}
}
