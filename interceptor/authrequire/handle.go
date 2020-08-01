package authrequire

import (
	"github.com/gin-gonic/gin"
	"github.com/sillyhatxu/gin-utils/v2/gincodes"
	"github.com/sillyhatxu/gin-utils/v2/jwtutils"
	"github.com/sillyhatxu/gin-utils/v2/response"
	"net/http"
)

type OptionService interface {
	GetJWTClient() *jwtutils.JWT
	GetTokenKey() string
	GetContextKey() string
	IsDebug() bool
	GetAuth(token string) (interface{}, error)
	GetAuthForDebug(ctx *gin.Context) (interface{}, error)
}

func AuthRequire(optionService OptionService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if optionService.IsDebug() {
			auth, err := optionService.GetAuthForDebug(ctx)
			if err != nil {
				ctx.Header("Content-Type", "application/json")
				ctx.JSON(http.StatusUnauthorized, response.NewError(gincodes.Unauthenticated, "Unauthenticated indicates the request does not have valid"))
				ctx.Abort()
				return
			}
			ctx.Set(optionService.GetContextKey(), auth)
			ctx.Next()
			return
		}
		token, err := ctx.Cookie(optionService.GetTokenKey())
		if err != nil {
			ctx.Header("Content-Type", "application/json")
			ctx.JSON(http.StatusUnauthorized, response.NewError(gincodes.Unauthorized, "Unauthenticated indicates the request does not have valid"))
			ctx.Abort()
			return
		}
		auth, err := optionService.GetAuth(token)
		if err != nil {
			ctx.Header("Content-Type", "application/json")
			ctx.JSON(http.StatusUnauthorized, response.NewError(gincodes.Unauthenticated, "Unauthenticated indicates the request does not have valid"))
			ctx.Abort()
			return
		}
		ctx.Set(optionService.GetContextKey(), auth)
		ctx.Next()
		return
	}
}
