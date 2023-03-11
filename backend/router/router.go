package router

import (
	"github.com/gin-gonic/gin"
	"github.com/li-zeyuan/CSEIdiom/backend/config"
	"github.com/li-zeyuan/CSEIdiom/backend/handler"
	"github.com/li-zeyuan/common/httptransfer"
)

func New() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Recovery())
	//engine.Use(gin.RecoveryWithWriter(gin.DefaultErrorWriter))
	//engine.Use(gzip.Gzip(gzip.DefaultCompression))

	// todo set request id to content
	//engine.Use(requestid.New())
	//engine.Use(timeout.New(timeout.WithHandler(func(c *gin.Context) {c.Next()}),timeout.WithTimeout(time.Duration(config.AppCfg.Timeout) * time.Second)))
	engine.POST("/api/login/wechat_login", handler.WechatLogin)

	engine.Use(httptransfer.JwtMiddleware(config.AppCfg.JwtSecret))
	engine.POST("/api/user/profile", handler.UserProfile)

	return engine
}
