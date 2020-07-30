package api

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	SecretKey  = "sillyhattest"
	TokenKey   = "SILLY-HAT-TOKEN"
	ContextKey = "SILLY-HAT-TOKEN"
)

type User struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}

func getAuthUser(ctx *gin.Context) *User {
	value, ok := ctx.Get(ContextKey)
	if !ok {
		return nil
	}
	user, ok := value.(User)
	if !ok {
		return nil
	}
	return &user
}
