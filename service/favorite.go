package service

import (
	"douyin/global"
	"douyin/model"
	"gorm.io/gorm"
)

type FavoriteService struct{}

func (FavoriteService) Action(param model.FavoriteActionParam) bool {
	// 判断ActionType，为1则点赞，为2则取消点赞，其他情况return false
	if param.ActionType == 1 {
		// 如果已经在点赞表，return true
		if err := global.Db.Where("user_id = ? and video_id = ?", param.UserId, param.VideoId).First(&model.Favorite{}).Error; err == nil {
			return true
		}

		// 点赞
		if err := global.Db.Transaction(func(tx *gorm.DB) error {
			if err := tx.Create(&model.Favorite{UserId: param.UserId, VideoId: param.VideoId}).Error; err != nil {
				return err
			}
			if err := tx.Model(&model.Video{}).Where("id = ?", param.VideoId).
				UpdateColumn("favorite_count", gorm.Expr("favorite_count + 1")).Error; err != nil {
				return err
			}
			return nil
		}); err != nil {
			return false
		}
		return true
	} else if param.ActionType == 2 {
		var favorite model.Favorite
		// 如果已不再点赞表，return true
		if err := global.Db.Where("user_id = ? and video_id = ?", param.UserId, param.VideoId).First(&favorite).Error; err != nil {
			return true
		}

		// 取消点赞
		if err := global.Db.Transaction(func(tx *gorm.DB) error {
			if err := tx.Delete(&favorite).Error; err != nil {
				return err
			}
			if err := tx.Model(&model.Video{}).Where("id = ?", param.VideoId).
				UpdateColumn("favorite_count", gorm.Expr("favorite_count - 1")).Error; err != nil {
				return err
			}
			return nil
		}); err != nil {
			return false
		}
		return true
	} else {
		return false
	}
}

func (FavoriteService) List(userId string) []model.VideoInfo {
	var videoIds []uint64
	global.Db.Model(&model.Favorite{}).Select("video_id").Where("user_id = ?", userId).Find(&videoIds)
	videoInfos := make([]model.VideoInfo, 0)
	for _, videoId := range videoIds {
		var video model.Video
		var author model.UserInfo

		// 查询视频详细信息
		global.Db.Model(&model.Video{}).Where("id = ?", videoId).First(&video)

		// 通过AuthorId查找视频作者相关信息
		global.Db.Model(&model.User{}).Where("id = ?", video.AuthorId).First(&author)

		var isFollowCount, isFavoriteCount int64
		// 查询是否已follow
		global.Db.Model(&model.Relation{}).Where("user_id = ? and follower_id = ?", author.Id, userId).Count(&isFollowCount)
		author.IsFollow = isFollowCount > 0

		// 查询是否已喜欢
		global.Db.Model(&model.Favorite{}).Where("user_id = ? and video_id = ?", userId, video.Id).Count(&isFavoriteCount)

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
