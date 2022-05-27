package model

// Relation 数据库Relation数据映射模型
type Relation struct {
	Id         uint64 `gorm:"column:id;autoIncrement;primaryKey"`
	UserId     uint64 `gorm:"column:user_id"`
	FollowerId uint64 `gorm:"column:follower_id"`
}

// FollowParam 关注操作参数
type FollowParam struct {
	UserId     uint64 `form:"user_id" json:"user_id"`
	ToUserId   uint64 `form:"to_user_id" json:"to_user_id"`
	ActionType int    `form:"action_type" json:"action_type"`
}

// UserInfoListResponse 获取用户列表
type UserInfoListResponse struct {
	Response
	UserInfoList []UserInfo `json:"user_list"`
}
