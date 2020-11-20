package logic

import (
	"BBS/dao/mysql"
	"BBS/models"
	"BBS/pkg/snowflake"
)

func CreatePost(p *models.Post) (err error){
	p.ID = int64(snowflake.GenID())
	return mysql.CreatePost(p)

}