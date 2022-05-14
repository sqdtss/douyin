package api

import (
	"douyin/model"
	"douyin/service"
	"douyin/status"
	"github.com/gin-gonic/gin"
	"net/http"
)

var userService service.UserService

// Register 新用户注册时提供用户名，密码，昵称即可，用户名需要保证唯一。创建成功后返回用户id和权限token
func Register(c *gin.Context) {
	// 获取参数
	var param model.RegisterAndLoginParam
	if err := c.ShouldBind(&param); err != nil {
		c.JSON(http.StatusOK, model.RegisterAndLoginResponse{
			Response: model.Response{
				StatusCode: status.RequestParamError,
				StatusMsg:  status.Msg(status.RequestParamError),
			},
		})
		return
	}

	// 注册
	if ok, userId := userService.Register(param); ok {
		token, err := userService.GenerateToken(userId)
		// 未成功生成token
		if err != nil {
			c.JSON(http.StatusOK, model.RegisterAndLoginResponse{
				Response: model.Response{
					StatusCode: status.GenerateTokenError,
					StatusMsg:  status.Msg(status.GenerateTokenError),
				},
			})
			return
		}

		// 注册成功
		c.JSON(http.StatusOK, model.RegisterAndLoginResponse{
			Response: model.Response{
				StatusCode: status.Success,
				StatusMsg:  status.Msg(status.Success),
			},
			UserId: userId,
			Token:  token,
		})
	} else {
		// 已经有用户名为param.Username的用户
		c.JSON(http.StatusOK, model.RegisterAndLoginResponse{
			Response: model.Response{
				StatusCode: status.UsernameHasExistedError,
				StatusMsg:  status.Msg(status.UsernameHasExistedError),
			},
		})
	}
}

// Login 通过用户名和密码进行登录，登录成功后返回用户id和权限token
func Login(c *gin.Context) {
	// 获取参数
	var param model.RegisterAndLoginParam
	if err := c.ShouldBind(&param); err != nil {
		c.JSON(http.StatusOK, model.RegisterAndLoginResponse{
			Response: model.Response{
				StatusCode: status.RequestParamError,
				StatusMsg:  status.Msg(status.RequestParamError),
			},
		})
		return
	}

	// 登录
	if ok, userId := userService.Login(param); ok {
		token, err := userService.GenerateToken(userId)
		// 未成功生成token
		if err != nil {
			c.JSON(http.StatusOK, model.RegisterAndLoginResponse{
				Response: model.Response{
					StatusCode: status.GenerateTokenError,
					StatusMsg:  status.Msg(status.GenerateTokenError),
				},
			})
			return
		}

		// 登录成功
		c.JSON(http.StatusOK, model.RegisterAndLoginResponse{
			Response: model.Response{
				StatusCode: status.Success,
				StatusMsg:  status.Msg(status.Success),
			},
			UserId: userId,
			Token:  token,
		})
	} else {
		// 登录失败，用户名不存在或密码错误
		c.JSON(http.StatusOK, model.RegisterAndLoginResponse{
			Response: model.Response{
				StatusCode: status.UserNotExistOrPasswordWrongError,
				StatusMsg:  status.Msg(status.UserNotExistOrPasswordWrongError),
			},
		})
	}
}

// UserInfo 获取登录用户的id、昵称，如果实现社交部分的功能，还会返回关注数和粉丝数
func UserInfo(c *gin.Context) {
	// 获取参数token
	token := c.Query("token")

	// 未获取到token
	if token == "" {
		c.JSON(http.StatusOK, model.UserInfoResponse{
			Response: model.Response{
				StatusCode: status.RequestParamError,
				StatusMsg:  status.Msg(status.RequestParamError),
			},
		})
		return
	}

	// 通过token获取id
	id, err := userService.GetIdByToken(token)
	if err != nil {
		c.JSON(http.StatusOK, model.UserInfoResponse{
			Response: model.Response{
				StatusCode: status.GetIdByTokenError,
				StatusMsg:  status.Msg(status.GetIdByTokenError),
			},
		})
		return
	}

	// 获取用户信息
	userInfo := userService.GetUserInfo(id)
	c.JSON(http.StatusOK, model.UserInfoResponse{
		Response: model.Response{
			StatusCode: status.Success,
			StatusMsg:  status.Msg(status.Success),
		},
		UserInfo: userInfo,
	})
}
