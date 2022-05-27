package api

import (
	"douyin/model"
	"douyin/service"
	"douyin/status"
	"github.com/gin-gonic/gin"
	"net/http"
)

var relationService service.RelationService

// RelationAction 关注操作
func RelationAction(c *gin.Context) {
	var param model.FollowParam
	if err := c.ShouldBind(&param); err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: status.RequestParamError,
			StatusMsg:  status.Msg(status.RequestParamError),
		})
		return
	}

	// 关注/取消关注
	if ok := relationService.Action(param); ok {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: status.Success,
			StatusMsg:  status.Msg(status.Success),
		})
		return
	}

	// 未知错误
	c.JSON(http.StatusOK, model.Response{
		StatusCode: status.UnknownError,
		StatusMsg:  status.Msg(status.UnknownError),
	})
}

// FollowList 关注列表
func FollowList(c *gin.Context) {
	// 获取参数user_id
	userId := c.Query("user_id")

	c.JSON(http.StatusOK, model.UserInfoListResponse{
		Response: model.Response{
			StatusCode: status.Success,
			StatusMsg:  status.Msg(status.Success),
		},
		UserInfoList: relationService.FollowList(userId),
	})
	return
}

// FollowerList 粉丝列表
func FollowerList(c *gin.Context) {
	// 获取参数user_id
	userId := c.Query("user_id")

	c.JSON(http.StatusOK, model.UserInfoListResponse{
		Response: model.Response{
			StatusCode: status.Success,
			StatusMsg:  status.Msg(status.Success),
		},
		UserInfoList: relationService.FollowerList(userId),
	})
}
