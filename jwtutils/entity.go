package jwtutils

import (
	"encoding/base64"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

const (
	DebugKey      = "debug"
	JwtEntity     = "JWT_ENTITY"
	TokenKey      = "SILLY-HAT-TOKEN"
	Authorization = "Authorization"
)

var Client *JWT

type JWT struct {
	initial   bool
	secretKey []byte
}

func (j JWT) SecretKey() string {
	return base64.URLEncoding.EncodeToString(j.secretKey)
}

func (j JWT) CreateToken(claims jwt.Claims) (string, error) {
	return createTokenStringHS512(j.secretKey, claims)
}

func (j JWT) ParseToken(token string, claims jwt.Claims) error {
	tokenEntity, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return j.secretKey, nil
	})
	if err != nil {
		return err
	}
	if !tokenEntity.Valid {
		return fmt.Errorf("token valid failed")
	}
	return nil
}

func InitialJWT(secretKey string) error {
	dec, err := base64.URLEncoding.DecodeString(secretKey)
	if err != nil {
		return err
	}
	Client = &JWT{
		initial:   true,
		secretKey: dec,
	}
	return nil
}
