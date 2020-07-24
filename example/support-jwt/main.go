package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sillyhatxu/gin-utils"
	"github.com/sillyhatxu/gin-utils/gincodes"
	"github.com/sillyhatxu/gin-utils/jwtutils"
	"github.com/sirupsen/logrus"
	"net"
	"net/http"
)

func main() {
	InitialAPI()
}

func InitialAPI() {
	err := jwtutils.InitialJWT("SILLYHAT")
	if err != nil {
		panic(err)
	}
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		panic(err)
	}
	logrus.Info("initial internal api start")
	router := SetupRouter()
	err = http.Serve(listener, router)
	if err != nil {
		panic(err)
	}
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/login", login)
	router.GET("/get", get)
	return router
}

type LoginDTO struct {
	LoginName string `json:"loginName"`
	Password  string `json:"password"`
}

type User struct {
	UserId   string `json:"userId"`
	UserName string `json:"userName"`
	Status   bool   `json:"status"`
}

func (u User) Valid() error {
	return nil
}

var TokenKey = "SILLY-HAT-TOKEN"

func get(ctx *gin.Context) {
	token, err := ctx.Cookie(TokenKey)
	if err != nil {
		ctx.JSON(http.StatusOK, ginutils.Errorf(gincodes.PermissionDenied, err))
		return
	}
	var u User
	err = jwtutils.Client.ParseToken(token, &u)
	if err != nil {
		ctx.JSON(http.StatusOK, ginutils.Errorf(gincodes.PermissionDenied, err))
		return
	}
	ctx.JSON(http.StatusOK, ginutils.Success(ginutils.Data(u)))
	return
}

func login(ctx *gin.Context) {
	var loginDTO *LoginDTO
	err := ctx.ShouldBindJSON(&loginDTO)
	if err != nil {
		ctx.JSON(http.StatusOK, ginutils.Errorf(gincodes.InvalidArgument, err))
		return
	}
	logrus.Infof("loginDTO : %#v", loginDTO)
	u := User{
		UserId:   "1",
		UserName: "test user",
		Status:   true,
	}
	token, err := jwtutils.Client.CreateToken(u)
	if err != nil {
		ctx.JSON(http.StatusOK, ginutils.Errorf(gincodes.PermissionDenied, err))
		return
	}
	ctx.SetCookie(TokenKey, token, 60*60*24, "/", "localhost", false, true)
	//ctx.SetCookie("SILLY-HAT-TOKEN", token, 60*60*24, "/", "http://localhost:8080", true, true)
	ctx.JSON(http.StatusOK, ginutils.Success(ginutils.Data(map[string]string{"token": token})))
	return
}
