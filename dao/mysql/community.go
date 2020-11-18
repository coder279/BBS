package mysql

import (
	"BBS/models"
	"database/sql"
	"errors"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := `select community_id, community_name from community`
	err = db.Select(&communityList, sqlStr)
	if err == sql.ErrNoRows {
		err = nil
		return
	}
	return
}

func GetCommunityDetailById(id int64)(*models.CommunityDetail,error){
	community := new(models.CommunityDetail)
	sqlStr := `select community_id,community_name,introduction,create_time from community where community_id = ? `
	err := db.Get(community,sqlStr,id)
	if err != nil {
		if err == sql.ErrNoRows{
			err = errors.New("无效的id")
		}
	}
	return community,err
}
