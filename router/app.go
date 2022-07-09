package router

import (
	"gin_gorm_object/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	// 配置路由
	r.GET("/ping", service.Ping)

	return r
}
