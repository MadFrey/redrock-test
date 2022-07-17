package main

import (
	_ "github.com/go-sql-driver/mysql"
	"redrock-test/controller"
	"redrock-test/dao"
	"redrock-test/util"
)

func main() {
	util.InitLogger()
	dns := "root:sjk123456@tcp(localhost:3306)/chess"
	util.SugarLogger.Debug("初始化数据库")
	err := dao.Init(dns)
	if err != nil {
		util.SugarLogger.Fatal("数据库初始化失败")
		return
	}
	util.SugarLogger.Debug("初始化路由")
	r := controller.CreateRouter()
	err = r.Run(":9090")
	if err != nil {
		util.SugarLogger.Fatal("路由初始化失败")
		return
	}
}
