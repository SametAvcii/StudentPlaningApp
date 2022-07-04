package handler

import (
	"fmt"
	model "github.com/SametAvcii/StudentPlaningApp/api/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
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
func GetUsersByID(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var users []*model.User
		id, _ := strconv.Atoi(c.Param("id"))
		if err := db.First(&users, id).Error; err != nil {
			return err
		}
		return c.JSON(http.StatusOK, users)
	}
}
func GetPlansByUserID(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var plans []*model.Plan
		id, _ := strconv.Atoi(c.Param("id"))
		if err := db.Limit(10).Select("*").Where(&model.Plan{UserId: uint(id)}).Find(&plans).Error; err != nil {
			return err
		}
		return c.JSON(http.StatusOK, plans)
	}
}
func GetPlansByUsername(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var plans []*model.Plan
		username, _ := strconv.Atoi(c.Param("username"))
		if err := db.Limit(10).Select("*").Where(&model.Plan{Username: string(username)}).Find(&plans).Error; err != nil {
			return err
		}
		return c.JSON(http.StatusOK, plans)
	}
}

func UpdateUsers(db *gorm.DB, newUserName, name, password, email string) echo.HandlerFunc {

	return func(c echo.Context) error {
		var users []*model.User
		username, _ := strconv.Atoi(c.Param("username"))
		if err := db.Model(&users).Where(&model.User{UserName: string(username)}).Updates(model.User{Name: name, UserName: newUserName, Password: password, Email: email}).Error; err != nil {
			return err
		}

		return c.JSON(http.StatusOK, users)
	}
}
func UpdatePlans(db *gorm.DB, content, state string, start_Date, finish_Date time.Time) echo.HandlerFunc {

	return func(c echo.Context) error {
		var plans []*model.Plan
		name, _ := strconv.Atoi(c.Param("name"))
		if err := db.Model(&plans).Where(&model.Plan{Name: string(name)}).Updates(model.Plan{Name: string(name), Content: content, State: state, Start_date: start_Date, Finish_date: finish_Date}).Error; err != nil {
			return err
		}
		return c.JSON(http.StatusOK, plans)
	}
}

func DeleteUsers(db *gorm.DB) echo.HandlerFunc {

	return func(c echo.Context) error {
		var users []*model.User
		username, _ := strconv.Atoi(c.Param("username"))
		if err := db.Model(&users).Where(&model.User{UserName: string(username)}).Delete(&users).Error; err != nil {
			return err
		}

		return c.JSON(http.StatusOK, fmt.Sprintf("user %d deleted", username))
	}
}

func DeletePLanByUserID(db *gorm.DB) echo.HandlerFunc {

	return func(c echo.Context) error {

		id, _ := strconv.Atoi(c.Param("id"))
		name, _ := strconv.Atoi(c.Param("name"))
		var plan []*model.Plan
		if err := db.Model(&plan).Where(&model.Plan{UserId: uint(id), Name: string(name)}).Delete(&plan).Error; err != nil {
			return err
		}
		return c.JSON(http.StatusOK, plan)
	}
}
func DeletePLanByUsername(db *gorm.DB) echo.HandlerFunc {

	return func(c echo.Context) error {
		var plan []*model.Plan
		username, _ := strconv.Atoi(c.Param("username"))
		name, _ := strconv.Atoi(c.Param("name"))
		if err := db.Model(&plan).Where(&model.Plan{Username: string(username), Name: string(name)}).Delete(&plan).Error; err != nil {
			return err
		}
		return c.JSON(http.StatusOK, plan)
	}
}
