package api

import (
	model "github.com/SametAvcii/StudentPlaningApp/api/model"
	cl "github.com/SametAvcii/StudentPlaningApp/database/client"
	"gorm.io/gorm"
)

func CreatePlan(cpr model.Plan, conn string) (tx *gorm.DB) {
	db := cl.Connection(conn)
	db.AutoMigrate(&model.Plan{})

	db.Create(&cpr)

	var createPlan model.Plan
	return db.First(&createPlan, 1)
}
