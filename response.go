package main

import (
	"fmt"
	"github.com/SametAvcii/StudentPlaningApp/api"
	model "github.com/SametAvcii/StudentPlaningApp/api/model"
	cl "github.com/SametAvcii/StudentPlaningApp/database/client"
	"github.com/SametAvcii/StudentPlaningApp/util"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"time"
)

func main() {
	db := cl.Connection(util.DbConn)
	e := echo.New()

	api.ListPLanByDay(db, e, time.Saturday, "SametAvciii")
	api.ListPLanByMonth(db, e, time.January, "SametAvciii")
	var cpr model.Plan
	cpr.Model = gorm.Model{}
	cpr.State = "Active"
	cpr.UserId = 1
	cpr.Username = "SametAvciii"
	cpr.Name = "Js"
	cpr.Content = "Task For Internship"
	cpr.Start_date = time.Date(2022, 1, 1, 1, 1, 1, 1, time.UTC)
	cpr.Finish_date = time.Date(2022, 1, 1, 1, 1, 1, 1, time.UTC)
	//fmt.Println(cpr.Start_date.Weekday())
	conn := util.DbConn
	fmt.Println(api.CreatePlan(cpr, conn))

}
