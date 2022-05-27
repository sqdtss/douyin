package initialize

import (
	"douyin/api"
	"douyin/global"
	"douyin/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Router() {

	engine := gin.Default()

	// 静态资源请求映射
	engine.Static("/public", global.Config.Upload.SavePath)

	// douyin apis
	douyin := engine.Group("/douyin")
	{
		// feed api
		douyin.GET("/feed/", api.Feed)

		// user注册/登录api
		douyin.POST("/user/register/", api.Register)
		douyin.POST("/user/login/", api.Login)

		// 以下使用Jwt鉴权
		douyin.Use(middleware.JwtAuth())

		/// 获取用户信息api
		douyin.GET("/user/", api.UserInfo)

		// publish apis
		publish := douyin.Group("/publish")
		{
			publish.POST("/action/", api.Publish)
			publish.GET("/list/", api.PublishList)
		}

		// favorite apis
		favorite := douyin.Group("/favorite")
		{
			favorite.POST("/action/", api.FavoriteAction)
			favorite.GET("/list/", api.FavoriteList)
		}

		// comment apis
		comment := douyin.Group("/comment")
		{
			comment.POST("/action/", api.CommentAction)
			comment.GET("/list/", api.CommentList)
		}

		// relation apis
		relation := douyin.Group("/relation")
		{
			relation.POST("/action/", api.RelationAction)
			relation.GET("/follow/list/", api.FollowList)
			relation.GET("/follower/list/", api.FollowerList)
		}
	}
	// 启动、监听端口
	post := fmt.Sprintf(":%s", global.Config.Server.Port)
	if err := engine.Run(post); err != nil {
		fmt.Printf("engine run error: %s", err)
	}
}
