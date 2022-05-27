package api

import (
	"douyin/model"
	"douyin/service"
	"douyin/status"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

var publishService service.PublishService

// Publish 登录用户选择视频上传
func Publish(c *gin.Context) {
	// 获取参数token, 注意：这个token是存在text中的
	token, ok := c.GetPostForm("token")

	// 未获取到token
	if !ok {
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

	// 读取data
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: status.LoadFileError,
			StatusMsg:  status.Msg(status.LoadFileError),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%d_%s", id, filename)
	saveFile := filepath.Join("./public/", finalName)

	// 获取参数title, 注意：这个title是存在text中的
	title, _ := c.GetPostForm("title")

	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: status.SaveUploadedFileError,
			StatusMsg:  status.Msg(status.SaveUploadedFileError),
		})
		return
	} else {
		if row := publishService.SaveVideo(id, saveFile, title); row == 0 {
			c.JSON(http.StatusOK, model.Response{
				StatusCode: status.SaveUploadedFileError,
				StatusMsg:  status.Msg(status.SaveUploadedFileError),
			})
			return
		}
	}

	// 上传成功
	c.JSON(http.StatusOK, model.Response{
		StatusCode: status.Success,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList 登录用户所发表的所有视频
// 注：publish_list和favorite_list的video_list为空时需要返回[]而不是nil
func PublishList(c *gin.Context) {
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

	// 返回发表的视频
	c.JSON(http.StatusOK, model.PublishResponse{
		Response: model.Response{
			StatusCode: status.Success,
			StatusMsg:  status.Msg(status.Success),
		},
		VideoInfoList: publishService.List(userId),
	})
}
