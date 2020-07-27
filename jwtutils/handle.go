package jwtutils

import (
	"github.com/gin-gonic/gin"
	"github.com/sillyhatxu/gin-utils/gincodes"
	"github.com/sillyhatxu/gin-utils/response"
	"net/http"
)

func AuthRequire(secret string, input interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		isDebug := ctx.GetHeader(DebugKey)
		if isDebug == "true" {

		}
		token, err := ctx.Cookie("SILLY-HAT-TOKEN")
		if err != nil {
			ctx.Header("Content-Type", "application/json")
			ctx.JSON(http.StatusUnauthorized, response.NewError(gincodes.Unauthorized, "You are not authorized to access this page"))
			ctx.Abort()
			return
		}
		err = ParseToken(token, secret, &input)
		if err != nil {
			ctx.Header("Content-Type", "application/json")
			ctx.JSON(http.StatusUnauthorized, response.NewError(gincodes.ServerError, "error parsing input"))
			ctx.Abort()
			return
		}

		ctx.Set(constants.JwtEntity, input)
		ctx.Next()
		//t := time.Now()
		//// Set example variable
		//ctx.Set("example", "12345")
		//// before request
		//ctx.Next()
		//// after request
		//latency := time.Since(t)
		//log.Print(latency)
		//// access the status we are sending
		//status := ctx.Writer.Status()
		//log.Println(status)
	}
}
