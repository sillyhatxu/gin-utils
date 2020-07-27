package jwtutils

import (
	"encoding/base64"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

type JWT struct {
	secretKey []byte
}

func New(secretKey string) (*JWT, error) {
	dec, err := base64.URLEncoding.DecodeString(secretKey)
	if err != nil {
		return nil, err
	}
	return &JWT{
		secretKey: dec,
	}, nil
}

func (j *JWT) SecretKey() string {
	return base64.URLEncoding.EncodeToString(j.secretKey)
}

func (j *JWT) CreateToken(claims jwt.Claims) (string, error) {
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
