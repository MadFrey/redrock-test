package model

import (
	"database/sql"
	"log"
	"redrock-test/dao"
	"redrock-test/util"
)

type User struct {
	Id       int    `gorm:"column:id"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
}

func QueryUserPwd(username string) string {
	sqlstr := "select password from user where username=?"
	row := dao.DB.QueryRow(sqlstr, username)
	pwd := ""
	err := row.Scan(&pwd)
	if err != nil {
		log.Println(err)
		return ""
	}
	return pwd
}

func InsertUser(DB *sql.DB, user User) (int64, error) {
	sqlstr := "insert into user(username,password) values (?,?)"
	result, err := DB.Exec(sqlstr, user.Username, user.Password)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}

func QueryIdWithUsername(username string) int {
	sqlstr := "select id from user where username=?"
	row := dao.DB.QueryRow(sqlstr, username)
	id := 0
	err := row.Scan(&id)
	if err != nil {
		util.SugarLogger.Error(err)
		return 0
	}
	return id
}
