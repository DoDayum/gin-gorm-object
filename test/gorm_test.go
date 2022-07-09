package test

import (
	"fmt"
	"gin_gorm_object/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestGormTest(t *testing.T) {
	dsn := "root:970617Baiyu!@tcp(localhost:3306)/gin_gorm_object?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
		return
	}

	data := make([]*models.Problem, 0)
	err = db.Debug().Find(&data).Error

	if err != nil {
		t.Fatal(err)
	}

	for _, datum := range data {
		fmt.Printf("Problem = %v \n", datum)
	}

}
