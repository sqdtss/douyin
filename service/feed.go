package service

import (
	"douyin/global"
	"douyin/model"
)

type FeedService struct{}

func (FeedService) Feed(userId uint64, latestTime int64) []model.VideoInfo {
	var videos []model.Video
	global.Db.Model(&model.Video{}).Where("create_time < ?", latestTime).Order("create_time DESC").Limit(30).Find(&videos)
	var videoInfos []model.VideoInfo

	for _, video := range videos {
		var author model.UserInfo
		// 通过AuthorId查找视频作者相关信息
		global.Db.Model(&model.User{}).Where("id = ?", video.AuthorId).First(&author)

		var isFollowCount, isFavoriteCount int64
		if userId > 0 { // 解析出了用户id
			// 查询是否已follow
			global.Db.Model(&model.Relation{}).Where("user_id = ? and follower_id = ?", author.Id, userId).Count(&isFollowCount)
			author.IsFollow = isFollowCount > 0

			// 查询是否已喜欢
			global.Db.Model(&model.Favorite{}).Where("user_id = ? and video_id = ?", userId, video.Id).Count(&isFavoriteCount)
		}

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
