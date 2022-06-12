package controller

import (
	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	r:=gin.Default()
	r.POST("/user/register",RegisterPost)
	r.POST("/user/login", LoginPost)
	r.GET("/game",Chess)
	return r
}