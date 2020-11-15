package mysql

import (
	"BBS/models"
	"crypto/md5"
	"encoding/hex"
)

const secret = "lichen.com"

func CheckUserExist(username string)(bool,error){
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	err := db.Get(&count,sqlStr,username)
	if err != nil {
		return false,err
	}
	return count > 0,nil
}

func QueryUserByUsername(){

}
//InsertUser 对用户数据执行入库操作
func InsertUser(user *models.User)(err error){
	//对密码进行加密
	password := encryptPassword(user.Password)
	//执行SQL语句入库
	sqlstr := "insert into user(user_id,username,password) values(?,?,?)"
	db.Exec(sqlstr,user.UserID,user.Username,password)
	return
}

func encryptPassword(oPassword string)string{
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
