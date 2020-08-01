package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sillyhatxu/gin-utils/v2/gincodes"
	"github.com/sillyhatxu/gin-utils/v2/jwtutils"
	"github.com/sillyhatxu/gin-utils/v2/response"
)

type Auth struct {
	JWTClient  *jwtutils.JWT
	TokenKey   string
	ContextKey string
	Debug      bool
}

func (auth *Auth) GetJWTClient() *jwtutils.JWT {
	return auth.JWTClient
}

func (auth *Auth) GetTokenKey() string {
	return auth.TokenKey
}

func (auth *Auth) GetContextKey() string {
	return auth.ContextKey
}

func (auth *Auth) IsDebug() bool {
	return auth.Debug
}

func (auth *Auth) GetAuth(token string) (interface{}, error) {
	var user User
	err := auth.JWTClient.ParseToken(token, &user)
	return user, err
}

func (auth *Auth) GetAuthForDebug(ctx *gin.Context) (interface{}, error) {
	userId := ctx.GetString("SILLY_HAT_USER_ID")
	if userId == "" {
		return nil, response.NewError(gincodes.Unauthenticated, "Unauthenticated indicates the request does not have valid")
	}
	return User{UserId: userId}, nil
}
