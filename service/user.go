package service

import (
	"douyin/global"
	"douyin/model"
	"github.com/golang-jwt/jwt"
	"strconv"
	"time"
)

type UserService struct{}

var SigningKey = []byte(global.Config.Jwt.SigningKey)

type Claims struct {
	Id uint64 `json:"id"`
	jwt.StandardClaims
}

// GenerateToken 生成Token
func (UserService) GenerateToken(id uint64) (string, error) {
	claims := Claims{id, jwt.StandardClaims{
		ExpiresAt: time.Now().Unix() + 60*60,
		Issuer:    strconv.FormatUint(id, 10),
	},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SigningKey)
}

// VerifyToken 验证Token
func (UserService) VerifyToken(tokenString string) error {
	_, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return SigningKey, nil
	})
	return err
}

// GetIdByToken 验证Token
func (UserService) GetIdByToken(tokenString string) (uint64, error) {
	var claims Claims
	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return SigningKey, nil
	})
	if err != nil {
		return 0, err
	}
	return claims.Id, nil
}

// Register 注册
func (UserService) Register(param model.RegisterAndLoginParam) (bool, uint64) {
	var count int64
	global.Db.Model(&model.User{}).Where("username = ?", param.Username).Count(&count)
	if count > 0 {
		return false, 0
	}
	user := model.User{
		Name:     param.Username,
		Username: param.Username,
		Password: param.Password,
	}
	if global.Db.Create(&user).RowsAffected == 1 {
		return true, user.Id
	}
	return false, 0
}

// Login 登录
func (UserService) Login(param model.RegisterAndLoginParam) (bool, uint64) {
	var user model.User
	row := global.Db.Where("username = ? and password = ?", param.Username, param.Password).First(&user).RowsAffected
	if row == 0 {
		// 不存在用户名为param.Username且密码为param.Password的条目
		return false, 0
	} else {
		return true, user.Id
	}
}

// GetUserInfo 获取用户信息
func (UserService) GetUserInfo(userId string) (userInfo model.UserInfo) {
	global.Db.Model(&model.User{}).Where("id = ?", userId).First(&userInfo)
	return
}
