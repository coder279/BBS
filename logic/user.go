package logic

import (
	"BBS/dao/mysql"
	"BBS/pkg/snowflake"
)
//存放业务逻辑代码
func Signup(){
	//1.判断用户存在不存在
	mysql.QueryUserByUsername()
	//2.生成UID
	snowflake.GenID()
	//3.密码加密
	mysql.InsertUser()
	//保存进数据库
}
