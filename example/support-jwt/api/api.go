package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sillyhatxu/gin-utils"
	"github.com/sillyhatxu/gin-utils/entity"
	"github.com/sillyhatxu/gin-utils/example/dto"
	"github.com/sillyhatxu/gin-utils/gincodes"
	"github.com/sillyhatxu/gin-utils/interceptor/authrequire"
	"github.com/sillyhatxu/gin-utils/jwtutils"
	"github.com/sillyhatxu/gin-utils/response"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

var JWTClient *jwtutils.JWT

func InitialAPI(port int) {
	router := SetupRouter()
	err := router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		logrus.Errorf("server down. %v", err)
		panic(err)
	}
}

func SetupRouter() *gin.Engine {
	router, err := ginutils.SetupRouter()
	if err != nil {
		panic(err)
	}
	client, err := jwtutils.New(SecretKey)
	if err != nil {
		logrus.Errorf("initial jwt error. %v", err)
		panic(err)
	}
	JWTClient = client
	auth := &Auth{
		JWTClient:  JWTClient,
		TokenKey:   TokenKey,
		ContextKey: ContextKey,
		Debug:      false,
	}
	loginGroup := router.Group("")
	{
		loginGroup.POST("/login", login)
	}
	userGroup := router.Group("/users").Use(authrequire.AuthRequire(auth))
	{
		userGroup.POST("", createUser)
		userGroup.PUT("/:id", modifyUser)
		userGroup.DELETE("/:id", deleteUser)
		userGroup.GET("/:id", getUserById)
		userGroup.GET("", queryUserByParams)
	}
	return router
}

func login(ctx *gin.Context) {
	token, err := JWTClient.CreateToken(User{UserId: fmt.Sprintf("%d", time.Now().Nanosecond())})
	if err != nil {
		ctx.JSON(http.StatusOK, response.Errorf(gincodes.ParamsValidateError, err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(entity.Data(token)))
	return
}

func createUser(ctx *gin.Context) {
	var user *dto.UserDTO
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusOK, response.Errorf(gincodes.ParamsValidateError, err))
		return
	}
	if user == nil {
		ctx.JSON(http.StatusOK, response.NewError(gincodes.ParamsValidateError, "body is nil"))
		return
	}
	user.SetUserId(uuid.New().String())
	if !user.Validate() {
		ctx.JSON(http.StatusOK, response.NewError(gincodes.ParamsValidateError, "user validate failed"))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(entity.Data(user)))
	return
}

func modifyUser(ctx *gin.Context) {
	id := ctx.Param("id")
	userName, age, status := "test-name", 18, true
	user := dto.UserDTO{
		UserId:   &id,
		UserName: &userName,
		Age:      &age,
		Status:   &status,
	}
	ctx.JSON(http.StatusOK, response.Success(entity.Data(user)))
	return
}

func deleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, response.Success(entity.Data(id)))
	return
}

func getUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	userName, age, status := "test-name", 18, true
	user := dto.UserDTO{
		UserId:   &id,
		UserName: &userName,
		Age:      &age,
		Status:   &status,
	}
	ctx.JSON(http.StatusOK, response.Success(entity.Data(user)))
	return
}

func queryUserByParams(ctx *gin.Context) {
	ids := ctx.QueryArray("ids")
	userName := ctx.Query("userName")
	status := ctx.Query("status")
	limit := ctx.DefaultQuery("limit", "20")
	offset := ctx.DefaultQuery("offset", "0")

	user := getAuthUser(ctx)
	if user == nil {
		ctx.JSON(http.StatusUnauthorized, response.NewError(gincodes.Unauthorized, "Unauthenticated indicates the request does not have valid"))
		return
	}
	params := map[string]interface{}{
		"ids":      ids,
		"userName": userName,
		"status":   status,
		"limit":    limit,
		"offset":   offset,
		"user":     user,
	}
	extra := map[string]interface{}{
		"total":   50,
		"isFirst": offset == "0",
		"isLast":  false,
	}
	ctx.JSON(http.StatusOK, response.Success(entity.Data(params), entity.Extra(extra)))
	return
}
