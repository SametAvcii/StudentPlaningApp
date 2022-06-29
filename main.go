package main

import (
	"fmt"
	api "github.com/SametAvcii/StudentPlaningApp/api"
	model "github.com/SametAvcii/StudentPlaningApp/api/model"
	util "github.com/SametAvcii/StudentPlaningApp/util"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

func main() {
	var cpr model.Plan
	cpr.Model = gorm.Model{}
	cpr.Name = "Js"
	cpr.Content = "Task For Internship"
	cpr.Start_date = "2022-01-01 00:00:01"
	cpr.Finish_date = "2022-01-02 00:00:01"

	conn := util.DbConn
	fmt.Println(api.CreatePlan(cpr, conn))

	var user model.User
	user.Name = "Samet Avci"
	user.Password = "asdxd"
	user.Email = "sametavc05@gmail.com"

	fmt.Println(api.CreateUser(user, conn))

}
