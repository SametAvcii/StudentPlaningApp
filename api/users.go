package api

import (
	model "github.com/SametAvcii/StudentPlaningApp/api/model"
	cl "github.com/SametAvcii/StudentPlaningApp/database/client"
	"gorm.io/gorm"
)

func CreateUser(user model.User, conn string) (tx *gorm.DB) {
	db := cl.Connection(conn)
	db.AutoMigrate(&model.User{})

	db.Create(&user)
	var createUser model.User
	return db.First(&createUser, 1)
}
