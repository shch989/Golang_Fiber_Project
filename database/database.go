package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error
	const MYSQL = "root:981104shch98!@tcp(127.0.0.1:3306)/go_fiber_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := MYSQL
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Db 연결에 실패하였습니다.")
	}
	fmt.Println("Connected to database")
}
