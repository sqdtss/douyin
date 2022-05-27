package middleware

import (
	"douyin/model"
	"douyin/service"
	"douyin/status"
	"github.com/gin-gonic/gin"
	"net/http"
)

var userService service.UserService

// JwtAuth JWT认证中间件
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取参数token
		token, ok := c.GetPostForm("token")
		if !ok {
			token = c.Query("token")
		}

		if token == "" {
			c.JSON(http.StatusOK, model.Response{
				StatusCode: status.RequestParamError,
				StatusMsg:  status.Msg(status.RequestParamError),
			})
			c.Abort()
			return
		}
		if err := userService.VerifyToken(token); err != nil {
			c.JSON(http.StatusOK, model.Response{
				StatusCode: status.TokenExpiredError,
				StatusMsg:  status.Msg(status.TokenExpiredError),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
