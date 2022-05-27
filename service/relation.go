package service

import (
	"douyin/global"
	"douyin/model"
	"fmt"
	"gorm.io/gorm"
)

type RelationService struct{}

// Action 关注/取消关注
func (RelationService) Action(param model.FollowParam) bool {
	if param.ActionType == 1 {
		// 如果已经关注，return true
		if err := global.Db.Where("user_id = ? and follower_id = ?", param.ToUserId, param.UserId).First(&model.Relation{}).Error; err == nil {
			fmt.Println(1111)
			return true
		}

		// 关注
		if err := global.Db.Transaction(func(tx *gorm.DB) error {
			// relation表新增relation
			if err := tx.Create(&model.Relation{UserId: param.ToUserId, FollowerId: param.UserId}).Error; err != nil {
				return err
			}

			// to_user follower_count + 1
			if err := tx.Model(&model.User{}).Where("id = ?", param.ToUserId).
				UpdateColumn("follower_count", gorm.Expr("follower_count + 1")).Error; err != nil {
				return err
			}

			// user follow_count + 1
			if err := tx.Model(&model.User{}).Where("id = ?", param.UserId).
				UpdateColumn("follow_count", gorm.Expr("follow_count + 1")).Error; err != nil {
				return err
			}
			return nil
		}); err != nil {
			return false
		}
		return true
	} else if param.ActionType == 2 {
		// 如果没有关注，return true
		var relation model.Relation
		if err := global.Db.Where("user_id = ? and follower_id = ?", param.ToUserId, param.UserId).First(&relation).Error; err != nil {
			return true
		}

		// 取消关注
		if err := global.Db.Transaction(func(tx *gorm.DB) error {
			// relation表删除relation
			if err := tx.Delete(&relation).Error; err != nil {
				return err
			}

			// to_user follower_count - 1
			if err := tx.Model(&model.User{}).Where("id = ?", param.ToUserId).
				UpdateColumn("follower_count", gorm.Expr("follower_count - 1")).Error; err != nil {
				return err
			}

			// user follow_count - 1
			if err := tx.Model(&model.User{}).Where("id = ?", param.UserId).
				UpdateColumn("follow_count", gorm.Expr("follow_count - 1")).Error; err != nil {
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

// FollowList 返回follow的列表
func (RelationService) FollowList(userId string) []model.UserInfo {
	// 查找follow的relation
	var relations []model.Relation
	global.Db.Model(&model.Relation{}).Where("follower_id = ?", userId).Find(&relations)
	fmt.Println(relations)
	var userInfos []model.UserInfo
	for _, relation := range relations {
		var userInfo model.UserInfo
		if global.Db.Model(&model.User{}).Where("id = ?", relation.UserId).First(&userInfo).RowsAffected > 0 {
			userInfo.IsFollow = true
			userInfos = append(userInfos, userInfo)
		}
	}
	return userInfos
}

// FollowerList 返回follower的列表
func (RelationService) FollowerList(userId string) []model.UserInfo {
	// 查找follower的relation
	var relations []model.Relation
	global.Db.Model(&model.Relation{}).Where("user_id = ?", userId).Find(&relations)

	var userInfos []model.UserInfo
	for _, relation := range relations {
		var userInfo model.UserInfo
		if global.Db.Model(&model.User{}).Where("id = ?", relation.FollowerId).First(&userInfo).RowsAffected > 0 {
			var isFollowCount int64
			// 查询是否已follow
			global.Db.Model(&model.Relation{}).Where("user_id = ? and follower_id = ?", relation.FollowerId, userId).Count(&isFollowCount)
			userInfo.IsFollow = isFollowCount > 0

			userInfos = append(userInfos, userInfo)
		}
	}
	return userInfos
}
