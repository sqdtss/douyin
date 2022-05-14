package api

import (
	"douyin/model"
	"douyin/service"
	"douyin/status"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var commentService service.CommentService

func CommentAction(c *gin.Context) {
	// 获取参数
	var param model.CommentActionParam
	if err := c.ShouldBind(&param); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, model.Response{
			StatusCode: status.RequestParamError,
			StatusMsg:  status.Msg(status.RequestParamError),
		})
		return
	}

	// 获取参数token
	token := c.Query("token")

	// 未获取到token
	if token == "" {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: status.RequestParamError,
			StatusMsg:  status.Msg(status.RequestParamError),
		})
		return
	}

	// 通过token获取id
	id, err := userService.GetIdByToken(token)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: status.GetIdByTokenError,
			StatusMsg:  status.Msg(status.GetIdByTokenError),
		})
		return
	}

	// 操作成功
	if ok := commentService.Action(id, param); ok {
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

// CommentList 返回视频所有评论会好点？？用不到userId吧
func CommentList(c *gin.Context) {
	// 获取参数token
	token := c.Query("token")

	// 未获取到token
	if token == "" {
		c.JSON(http.StatusOK, model.CommentListResponse{
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
		c.JSON(http.StatusOK, model.CommentListResponse{
			Response: model.Response{
				StatusCode: status.GetIdByTokenError,
				StatusMsg:  status.Msg(status.GetIdByTokenError),
			},
		})
		return
	}

	videoId := c.Query("video_id")
	if videoId == "" {
		c.JSON(http.StatusOK, model.CommentListResponse{
			Response: model.Response{
				StatusCode: status.RequestParamError,
				StatusMsg:  status.Msg(status.RequestParamError),
			},
		})
		return
	}

	c.JSON(http.StatusOK, model.CommentListResponse{
		Response: model.Response{
			StatusCode: status.Success,
			StatusMsg:  status.Msg(status.Success),
		},
		CommentInfoList: commentService.List(id, videoId),
	})
}
