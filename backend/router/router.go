package router

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(gin.RecoveryWithWriter(gin.DefaultErrorWriter))
	engine.Use(gzip.Gzip(gzip.DefaultCompression))

	return engine
}
