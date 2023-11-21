package routers

import (
	"net/http"
	"webDevScaffold/logger"
	"webDevScaffold/settings"

	"github.com/gin-gonic/gin"
)

// Setup 注册路由
func Setup(mode string) *gin.Engine {
	// 发布版本
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true)) // 使用zap自定义中间件

	// 路由 简单示例
	r.GET("/version", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ver": settings.Conf.Version,
		})
	})

	return r
}
