package router

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/li-zeyuan/CSEIdiom/backend/handler"
	"github.com/li-zeyuan/common/httptransfer"
)

func New() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(gin.RecoveryWithWriter(gin.DefaultErrorWriter))
	engine.Use(gzip.Gzip(gzip.DefaultCompression))
	engine.POST("/api/login/wechat_login", handler.WechatLogin)

	// todo
	engine.Use(httptransfer.JwtMiddleware())
	engine.POST("/api/user/profile", handler.UserProfile)

	return engine
}
