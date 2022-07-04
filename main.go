package main

import (
	"github.com/SametAvcii/StudentPlaningApp/api"
	"github.com/SametAvcii/StudentPlaningApp/api/handler"
	model "github.com/SametAvcii/StudentPlaningApp/api/model"
	cl "github.com/SametAvcii/StudentPlaningApp/database/client"
	util "github.com/SametAvcii/StudentPlaningApp/util"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
	"time"
)

func main() {
	var cpr model.Plan
	cpr.Model = gorm.Model{}
	cpr.State = "Active"
	cpr.UserId = 2
	cpr.Username = "SametAvcii"
	cpr.Name = "Js"
	cpr.Content = "Task For Internship"
	cpr.Start_date = time.Date(2022, 1, 1, 1, 1, 1, 1, time.UTC)
	cpr.Finish_date = time.Date(2022, 1, 1, 1, 1, 1, 1, time.UTC)
	conn := util.DbConn
	//fmt.Println(api.CreatePlan(cpr, conn))

	//plans := []int{1, 2, 3, 4, 5}

	var user model.User
	user.Name = "Samet Avci"
	hashPassword, _ := util.HashPassword("SametAvcii.")
	user.Password = hashPassword
	user.Email = "sametavc05@gmail.com1"
	user.UserName = "SametAvci"

	api.CreateUser(user, conn)

	// Echo instance
	db := cl.Connection(conn)
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes

	//e.POST("/user/create", api.CreateUser(user,db)
	e.GET("/user", handler.GetUsers(db))
	e.GET("/plan", handler.GetPlans(db))
	e.GET("/user/:id", handler.GetUsersByID(db))
	e.GET("/user/plans/:id", handler.GetPlansByUserID(db))
	e.GET("/user/plans/:username", handler.GetPlansByUsername(db))
	e.POST("/user/update/:username", handler.UpdateUsers(db, "SametAvcii", "Samet", "Avci", "sasada@gmail.com"))
	e.POST("/plan/update/:name", handler.UpdatePlans(db, "Samet", "Avci", cpr.Start_date, cpr.Finish_date))
	e.DELETE("/user/delete/:username", handler.DeleteUsers(db))
	e.DELETE("/plan/delete/:id/:name", handler.DeletePLanByUserID(db))
	e.DELETE("/plan/delete/:username/:name", handler.DeletePLanByUsername(db))

	e.Logger.Fatal(e.Start(":1323"))

}
