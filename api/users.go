package api

import (
	model "github.com/SametAvcii/StudentPlaningApp/api/model"
	cl "github.com/SametAvcii/StudentPlaningApp/database/client"
	"gorm.io/gorm"
)

func CreateUser(cpr model.User, conn string) (tx *gorm.DB) {
	db := cl.Connection(conn)
	db.AutoMigrate(&model.User{})

	db.Create(&cpr)

	var createPlan model.User
	return db.First(&createPlan, 1)
}
