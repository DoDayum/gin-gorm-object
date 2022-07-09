package main

import (
	"gin_gorm_object/router"
)

func main() {
	r := router.Router()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

/*
go get -u gorm.io/gorm && go get -u gorm.io/driver/sqlite

*/
