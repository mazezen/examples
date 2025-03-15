package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mazezen/itools"
	"github.com/mazezen/zenchat/common/sdk"
	"github.com/mazezen/zenchat/internel/handler"
	"github.com/mazezen/zenchat/internel/middle"
	"time"
)

func RunHttpServer() error {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "ip=${remote_ip} time=${time_rfc3339}, method=${method}, uri=${uri}, status=${status}, latency_human=${latency_human}\n",
		Output: itools.EchoLog,
	}))
	e.Use(middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{echo.POST, echo.GET, echo.DELETE},
			AllowCredentials: true,
			MaxAge:           int(time.Hour) * 24,
		}))

	e.Use(middle.RequestLog())
	e.Use(middleware.BodyDumpWithConfig(middle.BodyDumpConfig))

	userHandler := handler.NewUserHandler()
	// 注册 登录
	e.POST("/chat/register", userHandler.Register) // 注册
	e.POST("/chat/lgn_pwd", userHandler.LgnPwd)    // 密码登录

	// 用户模块
	ug := e.Group("chat/user", middle.JwtAuth())
	{
		ug.POST("/update", userHandler.Update)   // 更新用户信息
		ug.DELETE("/delete", userHandler.Delete) // 注销
		ug.GET("/list", userHandler.List)        // 用户列表
	}
	// 聊天
	chatHandler := handler.NewChatHandler
	e.GET("/chat/send/message", chatHandler.Chat, middle.JwtAuth()) // 发送消息
	// 聊天记录
	e.POST("/chat/msg", chatHandler.ChatMsg, middle.JwtAuth())

	// 好友关系模块
	relationGroup := e.Group("chat/relation", middle.JwtAuth())
	relationHandler := handler.NewUserRelationHandler()
	{
		relationGroup.GET("/list", relationHandler.RelationList)                             // 好友关系列表
		relationGroup.POST("/friend/add/by/name", relationHandler.FriendAddByName)           // 添加好友
		relationGroup.DELETE("/friend/remove/by/name", relationHandler.RemoveRelationByName) // 解除指定好友关系
	}

	// 群聊关系模块
	communityGroup := e.Group("/chat/community", middle.JwtAuth())
	communityHandler := handler.NewCommunityHandler()
	{
		communityGroup.POST("/create", communityHandler.CreateCommunity)   // 创建群聊
		communityGroup.POST("/update", communityHandler.UpdateCommunity)   // 更新群聊信息
		communityGroup.POST("/join", communityHandler.JoinCommunity)       // 加入群聊
		communityGroup.GET("/list/by/owner", communityHandler.ListByOwner) // 获取群列表
	}

	// 图片、语音文件上传模块
	uploadGroup := e.Group("/chat/upload", middle.JwtAuth())
	uploadHandler := handler.NewUploadHandler
	{
		uploadGroup.POST("/file", uploadHandler.Upload) // 图片、语音文件上传
	}

	return e.Start(sdk.GetConf().HttpPort)
}
