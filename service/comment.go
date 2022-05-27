package service

import (
	"douyin/global"
	"douyin/model"
	"gorm.io/gorm"
	"time"
)

type CommentService struct{}

func (CommentService) Action(param model.CommentActionParam) bool {
	if param.ActionType == 1 {
		// 评论
		if err := global.Db.Transaction(func(tx *gorm.DB) error {
			if err := tx.Create(&model.Comment{
				UserId:     param.UserId,
				VideoId:    param.VideoId,
				Content:    param.CommentText,
				CreateDate: time.Now().String(),
			}).Error; err != nil {
				return err
			}
			if err := tx.Model(&model.Video{}).Where("id = ?", param.VideoId).
				UpdateColumn("comment_count", gorm.Expr("comment_count + 1")).Error; err != nil {
				return err
			}
			return nil
		}); err != nil {
			return false
		}
		return true
	} else if param.ActionType == 2 {
		// 删除评论
		if err := global.Db.Transaction(func(tx *gorm.DB) error {
			if err := tx.Delete(&model.Comment{}, param.CommentId).Error; err != nil {
				return err
			}
			if err := tx.Model(&model.Video{}).Where("id = ?", param.VideoId).
				UpdateColumn("comment_count", gorm.Expr("comment_count - 1")).Error; err != nil {
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

func (CommentService) List(videoId string) []model.CommentInfo {
	var comments []model.Comment
	global.Db.Model(&model.Comment{}).Where("video_id = ?", videoId).Find(&comments)
	var commentInfos []model.CommentInfo

	for _, comment := range comments {
		var user model.UserInfo
		// 通过userId查找视频user相关信息
		global.Db.Model(&model.User{}).Where("id = ?", comment.UserId).First(&user)

		// 注意，此函数没有care这时作者有没有follow
		commentInfos = append(commentInfos, model.CommentInfo{
			Id:         comment.Id,
			User:       user,
			Content:    comment.Content,
			CreateDate: comment.CreateDate,
		})
	}
	return commentInfos
}
