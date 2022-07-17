/**
 * @Author: lrc
 * @Date: 2022/7/17-9:54
 * @Desc: 路由注册
 **/

package controller

import (
	"github.com/gin-gonic/gin"
	"redrock-test/middleware"
)

func CreateRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.GinLogger(), middleware.GinZapRecovery(true))
	r.POST("/user/register", RegisterPost)
	r.POST("/user/login", LoginPost)
	r.GET("/game", Chess)
	return r
}
