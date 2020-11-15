package logic

import (
	"BBS/dao/mysql"
	"BBS/models"
	"BBS/pkg/snowflake"
	"errors"
)
//存放业务逻辑代码
func Signup(p *models.ParamSignUp)(err error){
	//1.判断用户存在不存在
	exist,err := mysql.CheckUserExist(p.Username)
	if err != nil {
		return err
	}
	if(exist){
		return errors.New("用户已经存在")
	}
	//2.生成UID
	userId := snowflake.GenID()
	user := &models.User{
		UserID: userId,
		Username: p.Username,
		Password: p.Password,
	}
	//3.密码加密
	err = mysql.InsertUser(user)
	//保存进数据库
	if err != nil {
		return errors.New("插入数据库失败")
	}
	return
}
