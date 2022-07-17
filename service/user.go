/**
 * @Author: lrc
 * @Date: 2022/7/16-23:27
 * @Desc: 用户操作
 **/

package service

import (
	"redrock-test/dao"
	"redrock-test/model"
	"redrock-test/util"
)

func JudgeUserExist(username string, password string) (bool, int) {
	pwd := model.QueryUserPwd(username)
	id := model.QueryIdWithUsername(username)
	verify := util.PasswordVerify(password, pwd)
	return verify, id
}

func AddNewUserProcess(username string, password string) (int64, error) {
	// 用户数据
	hash, err := util.PasswordHash(password)
	if err != nil {
		util.SugarLogger.Error(err)
		return 0, err
	}
	user := model.User{Username: username, Password: hash}
	// 返回
	return model.InsertUser(dao.DB, user)
}
