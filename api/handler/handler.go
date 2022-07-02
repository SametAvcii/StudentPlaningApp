package handler

import (
	model "github.com/SametAvcii/StudentPlaningApp/api/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func GetUsers(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var users []*model.User
		if err := db.Find(&users).Error; err != nil {
			return err
		}
		return c.JSON(http.StatusOK, users)
	}
}
func GetPlans(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var plans []*model.Plan
		if err := db.Find(&plans).Error; err != nil {
			return err
		}
		return c.JSON(http.StatusOK, plans)
	}
}
func GetUsersByID(db *gorm.DB, id int) echo.HandlerFunc {
	return func(c echo.Context) error {
		var users []*model.User
		if err := db.First(&users, id).Error; err != nil {
			return err
		}
		return c.JSON(http.StatusOK, users)
	}
}
func GetPlansByUserID(db *gorm.DB, UserId uint) echo.HandlerFunc {
	return func(c echo.Context) error {
		var plans []*model.Plan
		if err := db.Limit(10).Select("Name").Where(&model.Plan{UserId: UserId}).Find(&plans).Error; err != nil {
			return err
		}
		return c.JSON(http.StatusOK, plans)
	}
}
func GetPlansByUsername(db *gorm.DB, username string) echo.HandlerFunc {
	return func(c echo.Context) error {
		var plans []*model.Plan
		if err := db.Limit(10).Select("Name").Where(&model.Plan{Username: username}).Find(&plans).Error; err != nil {
			return err
		}
		return c.JSON(http.StatusOK, plans)
	}
}

func UpdateUsers(db *gorm.DB, username, newUserName, name, password, email string) echo.HandlerFunc {

	return func(c echo.Context) error {
		var users []*model.User
		if err := db.Model(&users).Where(&model.User{UserName: username}).Updates(model.User{Name: name, UserName: newUserName, Password: password, Email: email}).Error; err != nil {
			return err
		}

		return c.JSON(http.StatusOK, users)
	}
}
func UpdatePlans(db *gorm.DB, name, content, state string, start_Date, finish_Date time.Time) echo.HandlerFunc {

	return func(c echo.Context) error {
		var plans []*model.Plan
		if err := db.Model(&plans).Where(&model.Plan{Name: name}).Updates(model.Plan{Name: name, Content: content, State: state, Start_date: start_Date, Finish_date: finish_Date}).Error; err != nil {
			return err
		}
		return c.JSON(http.StatusOK, plans)
	}
}

func DeleteUsers(db *gorm.DB, username string) echo.HandlerFunc {

	return func(c echo.Context) error {
		var users []*model.User
		if err := db.Model(&users).Where(&model.User{UserName: username}).Delete(&users).Error; err != nil {
			return err
		}

		return c.JSON(http.StatusOK, users)
	}
}

func DeletePLanByUserID(db *gorm.DB, plan_name string, user_id uint) echo.HandlerFunc {

	return func(c echo.Context) error {
		var plan []*model.Plan
		if err := db.Model(&plan).Where(&model.Plan{UserId: user_id, Name: plan_name}).Delete(&plan).Error; err != nil {
			return err
		}
		return c.JSON(http.StatusOK, plan)
	}
}
func DeletePLanByUsername(db *gorm.DB, plan_name string, username string) echo.HandlerFunc {

	return func(c echo.Context) error {
		var plan []*model.Plan
		if err := db.Model(&plan).Where(&model.Plan{Username: username, Name: plan_name}).Delete(&plan).Error; err != nil {
			return err
		}
		return c.JSON(http.StatusOK, plan)
	}
}
