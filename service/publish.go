package service

import (
	"douyin/global"
	"douyin/model"
	"fmt"
	"net"
	"strings"
	"time"
)

type PublishService struct{}

func (PublishService) SaveVideo(authorId uint64, playUrl, title string) int64 {
	// 获取服务器的对外ip
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		return 0
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip := strings.Split(localAddr.String(), ":")[0]

	return global.Db.Create(&model.Video{
		AuthorId: authorId,
		PlayUrl:  fmt.Sprintf("%s://%s:%s/%s", global.Config.Server.Protocol, ip, global.Config.Server.Port, playUrl),
		// 封面直接设置为了默认封面
		CoverUrl:   "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		CreateTime: time.Now().Unix(),
		Title:      title,
	}).RowsAffected
}

func (PublishService) List(id string) []model.VideoInfo {
	var videos []model.Video
	global.Db.Model(&model.Video{}).Where("author_id = ?", id).Find(&videos)
	videoInfos := make([]model.VideoInfo, 0)
	for _, video := range videos {
		var author model.UserInfo
		// 通过AuthorId查找视频作者相关信息
		global.Db.Model(&model.User{}).Where("id = ?", video.AuthorId).First(&author)

		var isFollowCount, isFavoriteCount int64
		// 查询是否已follow
		global.Db.Model(&model.Relation{}).Where("user_id = ? and follower_id = ?", author.Id, id).Count(&isFollowCount)
		author.IsFollow = isFollowCount > 0

		// 查询是否已喜欢
		global.Db.Model(&model.Favorite{}).Where("user_id = ? and video_id = ?", id, video.Id).Count(&isFavoriteCount)

		videoInfos = append(videoInfos, model.VideoInfo{
			Id:            video.Id,
			Author:        author,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    isFavoriteCount > 0,
			Title:         video.Title,
		})
	}
	return videoInfos
}
