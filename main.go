package main

import (
	"awesomeProject/controller"
	"awesomeProject/dao"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	dns := "root:123456@tcp(8.130.103.141:3306)/chess"
	err := dao.Init(dns)
	if err != nil {
		log.Println(err)
		return
	}
	r := controller.CreateRouter()
	err = r.Run(":9090")
	if err != nil {
		log.Println(err)
		return
	}

}
