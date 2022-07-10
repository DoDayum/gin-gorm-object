package service

import (
	"gin_gorm_object/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetParamList
// @Tags service/公共方法
// @Summary 问题列表
// @Param page query int false "请输入当前页，默认第一页"
// @Param size query int false "size"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /problem [get]
func GetParamList(c *gin.Context) {
	models.GetParamList()
	c.String(http.StatusOK, "get param list")
}
