package models

import "gorm.io/gorm"

type Submit struct {
	gorm.Model
	Identity        string `gorm:"column:identity;type:varchar(255);" json:"identity"`
	ProblemIdentity string `gorm:"column:problem_identity;type:varchar(255);" json:"problem_identity"`
	UserIdentity    string `gorm:"column:user_identity;type:varchar(36);" json:"user_identity"`
	Path            string `gorm:"column:path;type:varchar(255);" json:"path"`
	Status          string `gorm:"column:status;type:tinyint;" json:"status"`
}

func (tableName *Submit) TableName() string {
	return "submit"
}
