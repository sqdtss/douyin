package api

import (
	"douyin/model"
	"douyin/service"
	"douyin/status"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var commentService service.CommentService

func CommentAction(c *gin.Context) {
	// 获取参数
	var param model.CommentActionParam
	if err := c.ShouldBind(&param); err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: status.RequestParamError,
			StatusMsg:  status.Msg(status.RequestParamError),
		})
		return
	}

	// 操作成功
	if ok := commentService.Action(param); ok {
		c.JSON(http.StatusOK, model.CommentListResponse{
			Response: model.Response{
				StatusCode: status.Success,
				StatusMsg:  status.Msg(status.Success),
			},
			CommentInfoList: commentService.List(strconv.FormatUint(param.VideoId, 10)),
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

// CommentList 返回视频所有评论
func CommentList(c *gin.Context) {
	// 获取参数video_id
	videoId := c.Query("video_id")
	if videoId == "" {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: status.RequestParamError,
			StatusMsg:  status.Msg(status.RequestParamError),
		})
		return
	}

	c.JSON(http.StatusOK, model.CommentListResponse{
		Response: model.Response{
			StatusCode: status.Success,
			StatusMsg:  status.Msg(status.Success),
		},
		CommentInfoList: commentService.List(videoId),
	})
}
