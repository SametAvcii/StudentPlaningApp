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
	cpr.Username = "SametAvciii"
	cpr.Name = "Js"
	cpr.Content = "Task For Internship"
	cpr.Start_date = time.Date(2022, 1, 1, 1, 1, 1, 1, time.UTC)
	cpr.Finish_date = time.Date(2022, 1, 1, 1, 1, 1, 1, time.UTC)
	conn := util.DbConn
	//fmt.Println(api.CreatePlan(cpr, conn))

	//plans := []int{1, 2, 3, 4, 5}

	var user model.User
	user.Name = "Samet Avci"
	user.Password = "asdxd"
	user.Email = "sametavc05@gmail.com1"
	user.UserName = "SametAvciii"
	//user.Plans = plans

	api.CreateUser(user, conn)

	// Echo instance
	db := cl.Connection(util.DbConn)
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/user", handler.GetUsers(db))
	e.GET("/plan", handler.GetPlans(db))
	e.GET("/user/id", handler.GetUsersByID(db, 1))
	e.GET("/user/plans/id", handler.GetPlansByUserID(db, 2))
	e.GET("/user/plans/username", handler.GetPlansByUsername(db, "SametAvciii"))
	e.POST("/user/update", handler.UpdateUsers(db, "SametAvcii", "Samet", "Avci", "sdadasda", "sasada@gmail.com"))
	e.POST("/plan/update", handler.UpdatePlans(db, "SametAvcii", "Samet", "Avci", cpr.Start_date, cpr.Finish_date))
	e.DELETE("/user/delete/", handler.DeleteUsers(db, "User_Name"))
	e.DELETE("/plan/delete/", handler.DeletePLanByUserID(db, "PLan_name", 1))
	e.Logger.Fatal(e.Start(":1323"))

}
