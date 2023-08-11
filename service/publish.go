package service

import (
	"fmt"
	"simple-demo/model"
)

func GetPublishListByUserID(userid int64) []model.Video {
	baseDBModel := model.BaseDBModel{}
	db := baseDBModel.Connect()
	if db == nil {
		fmt.Println("连接失败")
	}
	// 从数据库中查询该ID的用户发布过的视频
	var videos []model.Video
	db.Preload("Author").Where("user_id = ?", userid).Find(&videos)
	for i := range videos {
		var author model.User
		db.First(&author, 2)
		videos[i].Author = author
	}
	return videos
}
