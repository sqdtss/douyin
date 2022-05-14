package model

// Favorite 数据库Favorite数据映射模型
type Favorite struct {
	Id      uint64 `gorm:"column:id;autoIncrement;primaryKey"`
	UserId  uint64 `gorm:"column:user_id"`
	VideoId uint64 `gorm:"column:video_id"`
}

// FavoriteActionParam 登录用户对视频点赞和取消点赞操作传递的参数
type FavoriteActionParam struct {
	VideoId    uint64 `form:"video_id" json:"video_id"`
	ActionType int    `form:"action_type" json:"action_type"`
}

// FavoriteListResponse 登录用户的点赞列表
type FavoriteListResponse struct {
	Response
	VideoInfoList []VideoInfo `json:"video_list"`
}
