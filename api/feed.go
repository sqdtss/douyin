package api

import (
	"douyin/model"
	"douyin/service"
	"douyin/status"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

var feedService service.FeedService

// Feed 无需登录，返回按投稿时间倒序的视频列表，视频数由服务端控制，单次最多30个。
func Feed(c *gin.Context) {
	var latestTime int64
	if latestTimeParam := c.Query("latest_time"); latestTimeParam != "" {
		// 解析为int64
		var err error
		latestTime, err = strconv.ParseInt(latestTimeParam, 10, 64)
		// 如果有err，表明传递过来的参数有误，返回RequestParamError
		if err != nil {
			c.JSON(http.StatusOK, model.Response{
				StatusCode: status.RequestParamError,
				StatusMsg:  status.Msg(status.RequestParamError),
			})
			return
		}
	} else {
		// latestTimeParam为空则表示未接收到latest_time参数，设置为nowtime
		latestTime = time.Now().Unix()
	}

	// 获取token
	token := c.Query("token")

	// 通过token获取id, 生成失败时id为0，数据库中无id为0的数据
	id, _ := userService.GetIdByToken(token)

	// 如果返回长度大于0则表示有视频，返回给用户
	if videoInfos := feedService.Feed(id, latestTime); len(videoInfos) > 0 {
		c.JSON(http.StatusOK, model.FeedResponse{
			Response: model.Response{
				StatusCode: status.Success,
				StatusMsg:  status.Msg(status.Success),
			},
			VideoInfoList: videoInfos,
			NextTime:      time.Now().Unix(),
		})
		return
	}

	// 其他情况，返回UnknownError
	c.JSON(http.StatusOK, model.FeedResponse{
		Response: model.Response{
			StatusCode: status.UnknownError,
			StatusMsg:  status.Msg(status.UnknownError),
		},
	})
}
