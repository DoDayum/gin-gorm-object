package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Identity string `gorm:"column:identity;type:varchar(36);" json:"identity"`
	Name     string `gorm:"column:name;type:varchar(255);" json:"identity"`
	ParentId string `gorm:"column:parent_id;type:varchar(255);" json:"parent_id"`
}

func (tableName *Category) TableName() string {
	return "category"
}
