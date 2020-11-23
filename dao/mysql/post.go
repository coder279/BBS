package mysql

import (
	"BBS/models"
)

func CreatePost(p *models.Post)(err error){
	sqlStr := `insert into post(post_id,title,content,author_id,community_id) values(?,?,?,?,?)`
	_,err = db.Exec(sqlStr,p.ID,p.Title,p.Content,p.AuthorID,p.CommunityID)
	return
}

func GetPostById(pid int64)(data *models.Post,err error){
	sqlStr := `select post_id,title,content,author_id,community_id,create_time from post where post_id =?`
	data = new (models.Post)
	err = db.Get(data,sqlStr,pid)
	return
}

func GetPostList(offset,limit int64) (data []*models.Post,err error){
	sqlStr := `select post_id,title,content,author_id,community_id,create_time from post limit ?,?`
	data = make([]*models.Post,0,2)
	err = db.Select(&data,sqlStr,offset,limit)
	return
}