package client

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func Connection(conn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
