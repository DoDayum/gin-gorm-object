package models

import "gorm.io/gorm"

type Problem struct {
	gorm.Model
	Identity   string `gorm:"column:identity;type:varchar(255);" json:"identity"`
	CategoryId string `gorm:"column:category_id;type:varchar(255);" json:"category_id"`
	Title      string `gorm:"column:title;type:varchar(255);" json:"title"`
	Content    string `gorm:"column:content;type:text;" json:"content"`
	MaxMem     string `gorm:"column:max_mem;type:text;" json:"max_mem"`
	MaxRuntime string `gorm:"column:max_runtime;type:text;" json:"max_runtime"`
}

func (table *Problem) TableName() string {
	return "problem"
}
