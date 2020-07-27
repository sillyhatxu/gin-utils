package authrequire

import (
	"github.com/gin-gonic/gin"
	"github.com/sillyhatxu/gin-utils/gincodes"
	"github.com/sillyhatxu/gin-utils/response"
	"net/http"
)

func AuthRequire(input interface{}, opts ...Option) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		config := &Config{
			JWTClient:  nil,
			TokenKey:   "",
			ContextKey: "",
			IsDebug:    false,
			DebugInput: nil,
		}
		for _, opt := range opts {
			opt(config)
		}
		if config.IsDebug {
			input = config.DebugInput
			ctx.Set(config.ContextKey, input)
			ctx.Next()
			return
		}
		token, err := ctx.Cookie(config.TokenKey)
		if err != nil {
			ctx.Header("Content-Type", "application/json")
			ctx.JSON(http.StatusUnauthorized, response.NewError(gincodes.Unauthorized, "You are not authorized to access this page"))
			ctx.Abort()
			return
		}
		err = config.JWTClient.ParseToken(token, &input)
		if err != nil {
			ctx.Header("Content-Type", "application/json")
			ctx.JSON(http.StatusUnauthorized, response.NewError(gincodes.ServerError, "error parsing input"))
			ctx.Abort()
			return
		}
		ctx.Set(config.ContextKey, input)
		ctx.Next()
	}
}
