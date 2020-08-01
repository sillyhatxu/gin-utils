package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sillyhatxu/gin-utils/v2"
	"github.com/sillyhatxu/gin-utils/v2/entity"
	"github.com/sillyhatxu/gin-utils/v2/example/dto"
	"github.com/sillyhatxu/gin-utils/v2/gincodes"
	"github.com/sillyhatxu/gin-utils/v2/response"
	"github.com/sirupsen/logrus"
	"net/http"
)

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

	//authorized.Use(attributes.RequireLogin())
	//home := router.Group("/home")
	//home.Use(requireLogin)
	//home.GET("", handleHome)
	userGroup := router.Group("/users")
	{
		userGroup.POST("", createUser)
		userGroup.PUT("/:id", modifyUser)
		userGroup.DELETE("", deleteUser)
		userGroup.GET("/:id", getUserById)
		userGroup.GET("", queryUserByParams)
	}
	return router
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

	params := map[string]interface{}{
		"ids":      ids,
		"userName": userName,
		"status":   status,
		"limit":    limit,
		"offset":   offset,
	}
	extra := map[string]interface{}{
		"total":   50,
		"isFirst": offset == "0",
		"isLast":  false,
	}
	ctx.JSON(http.StatusOK, response.Success(entity.Data(params), entity.Extra(extra)))
	return
}
