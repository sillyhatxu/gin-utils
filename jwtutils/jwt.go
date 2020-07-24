package jwtutils

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

const (
	HS256 = "HS256"
	HS384 = "HS384"
	HS512 = "HS512"
)

func createTokenString(secretKey []byte, signingKey string, claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.GetSigningMethod(signingKey), claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func createTokenStringHS256(secretKey []byte, claims jwt.Claims) (string, error) {
	return createTokenString(secretKey, HS256, claims)
}

func createTokenStringHS384(secretKey []byte, claims jwt.Claims) (string, error) {
	return createTokenString(secretKey, HS384, claims)
}

func createTokenStringHS512(secretKey []byte, claims jwt.Claims) (string, error) {
	return createTokenString(secretKey, HS512, claims)
}

func parseToken(token string, secretKey []byte, input interface{}) error {
	standardClaims := jwt.StandardClaims{}
	tokenEntity, err := jwt.ParseWithClaims(token, &standardClaims, func(token *jwt.Token) (interface{}, error) {
		//dec, err := base64.URLEncoding.DecodeString(secretKey)
		//if err != nil {
		//	return nil, err
		//}
		return secretKey, nil
	})
	if err != nil {
		return err
	}
	if !tokenEntity.Valid {
		return fmt.Errorf("token valid failed")
	}
	err = json.Unmarshal([]byte(standardClaims.Subject), &input)
	if err != nil {
		return err
	}
	return nil
}
