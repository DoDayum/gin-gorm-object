package router

import (
	_ "gin_gorm_object/docs" // 当前项目下doc
	"gin_gorm_object/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()

	// swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 配置路由
	r.GET("/ping", service.Ping)
	r.GET("/problem", service.GetParamList)

	return r
}
