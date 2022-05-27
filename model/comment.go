package model

// Comment 数据库Comment数据映射模型
type Comment struct {
	Id         uint64 `gorm:"column:id;autoIncrement;primaryKey"`
	UserId     uint64 `gorm:"column:user_id"`
	VideoId    uint64 `gorm:"column:video_id"`
	Content    string `gorm:"column:content"`
	CreateDate string `gorm:"column:create_date"`
}

// CommentActionParam 登录用户对视频评论操作传递的参数
type CommentActionParam struct {
	UserId      uint64 `form:"user_id" json:"user_id"`
	VideoId     uint64 `form:"video_id" json:"video_id"`
	ActionType  int    `form:"action_type" json:"action_type"`
	CommentText string `form:"comment_text" json:"comment_text"`
	CommentId   string `form:"comment_id" json:"comment_id"`
}

type CommentInfo struct {
	Id         uint64   `json:"id"`
	User       UserInfo `json:"user"`
	Content    string   `json:"content"`
	CreateDate string   `json:"create_date"`
}

// CommentListResponse 评论列表
type CommentListResponse struct {
	Response
	CommentInfoList []CommentInfo `json:"comment_list"`
}
