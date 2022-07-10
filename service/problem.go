package service

import (
	"gin_gorm_object/define"
	"gin_gorm_object/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// GetParamList
// @Tags service/公共方法
// @Summary 问题列表
// @Param page query int false "请输入当前页，默认第一页"
// @Param size query int false "size"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /problem [get]
func GetParamList(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	size, err := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	if err != nil {
		log.Println("get problem page error")
		return
	}
	keyWord := c.Query("keyWord")
	tx := models.GetParamList(keyWord)
	date := make([]*models.Problem, 0)
	page = (page - 1) * size
	var count int64
	err = tx.Debug().Omit("content").Count(&count).Offset(page).Limit(size).Find(&date).Error
	if err != nil {
		log.Println("get problem list error")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"date":  date,
		"count": count,
	})
}
