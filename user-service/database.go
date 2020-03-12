package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	pb "shippy/user-service/proto/user"
)

// 创建数据库连接
func ConnectDB(driver, dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(driver, dsn)

	if db != nil {
		log.Println("db connected")
		db.AutoMigrate(&pb.User{})
	}


	return db, err
}