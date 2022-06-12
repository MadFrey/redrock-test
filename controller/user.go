package controller

import (
	"awesomeProject/service"
	"awesomeProject/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterPost(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	rePassword := c.PostForm("rePassword")
	if username == "" {
		util.PrintInfo(c, "用户名不能为空！", 100)
		return
	}
	if len(username) > 20 {
		util.PrintInfo(c, "用户名长度超出限制！", 100)
		return
	}

	if len(password) > 16 {
		util.PrintInfo(c, "密码长度超出限制！", 100)
		return
	} else if len(password) < 8 {
		util.PrintInfo(c, "密码太短！", 100)
		return
	}
	if password != rePassword {
		util.PrintInfo(c, "两次密码不一致！", 100)
		return
	}

	flag, _ := service.JudgeUserExist(username, password)
	if flag {
		util.PrintInfo(c, "用户名已存在！", 100)
		return
	}
	_, err := service.AddNewUserProcess(username, password)
	if err != nil {
		util.PrintInfo(c, "注册失败！", 805)
		return
	}
	util.PrintInfo(c, "注册成功！请登录！", 0)
}

func LoginPost(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	flag, id := service.JudgeUserExist(username, password)
	if flag {
		tokenString, refreshTokenString := service.CreateToken(username)
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "登录成功，欢迎进入！",
			"data": gin.H{
				"uid":          id,
				"token":        tokenString,
				"refreshToken": refreshTokenString,
			},
		})
	} else {
		util.PrintInfo(c, "用户名或密码错误！", 1)
	}

}
