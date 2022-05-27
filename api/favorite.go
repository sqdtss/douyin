package api

import (
	"douyin/model"
	"douyin/service"
	"douyin/status"
	"github.com/gin-gonic/gin"
	"net/http"
)

var favoriteService service.FavoriteService

// FavoriteAction 登录用户对视频的点赞和取消点赞操作
func FavoriteAction(c *gin.Context) {
	// 获取参数
	var param model.FavoriteActionParam
	if err := c.ShouldBind(&param); err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: status.RequestParamError,
			StatusMsg:  status.Msg(status.RequestParamError),
		})
		return
	}

	// 操作成功
	if ok := favoriteService.Action(param); ok {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: status.Success,
			StatusMsg:  status.Msg(status.Success),
		})
		return
	}

	// 其他情况，发生未知错误
	c.JSON(http.StatusOK, model.Response{
		StatusCode: status.UnknownError,
		StatusMsg:  status.Msg(status.UnknownError),
	})
	return
}

// FavoriteList 登录用户的所有点赞视频
// 注：favorite_list和publish_list的video_list为空时需要返回[]而不是nil
func FavoriteList(c *gin.Context) {
	// 获取参数user_id
	userId := c.Query("user_id")

	// 未获取到user_id
	if userId == "" {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: status.RequestParamError,
			StatusMsg:  status.Msg(status.RequestParamError),
		})
		return
	}

	// 将用户喜欢的视频info返回
	c.JSON(http.StatusOK, model.FavoriteListResponse{
		Response: model.Response{
			StatusCode: status.Success,
			StatusMsg:  status.Msg(status.Success),
		},
		VideoInfoList: favoriteService.List(userId),
	})
}
