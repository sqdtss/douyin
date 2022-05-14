package model

// User 数据库User数据映射模型
type User struct {
	Id            uint64 `gorm:"column:id;autoIncrement;primaryKey"`
	Name          string `gorm:"column:name"`
	Username      string `gorm:"column:username"`
	Password      string `gorm:"column:password"`
	FollowCount   uint64 `gorm:"column:follow_count"`
	FollowerCount uint64 `gorm:"column:follower_count"`
	IsFollow      bool   `gorm:"column:is_follow"`
}

// RegisterAndLoginParam 注册/登录参数
type RegisterAndLoginParam struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

// RegisterAndLoginResponse 返回用户id和权限token
type RegisterAndLoginResponse struct {
	Response
	UserId uint64 `json:"id"`
	Token  string `json:"token"`
}

// UserInfo 返回用户信息
type UserInfo struct {
	Id            uint64 `json:"id"`
	Name          string `json:"name"`
	FollowCount   uint64 `json:"follow_count"`
	FollowerCount uint64 `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

// UserInfoResponse 获取登录用户的id, 昵称, 关注数, 粉丝数
type UserInfoResponse struct {
	Response
	UserInfo UserInfo `json:"user"`
}
