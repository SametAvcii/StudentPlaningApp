package api

import (
	"encoding/json"
	"fmt"
	model "github.com/SametAvcii/StudentPlaningApp/api/model"
	cl "github.com/SametAvcii/StudentPlaningApp/database/client"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
)

func CreatePlan(cpr model.Plan, conn string) (tx *gorm.DB) {
	db := cl.Connection(conn)
	db.AutoMigrate(&model.Plan{})

	db.Create(&cpr)

	var createPlan model.Plan
	if CheckPLan(createPlan) {
		return db.First(&createPlan, 1)
	}
	fmt.Println("Plan not created because you have already plan")
	return nil
}
func CheckPLan(cpr model.Plan) bool {
	var plans []model.Plan
	Resp := GetResponseResult("http://localhost:1323/user/plans")
	err := json.Unmarshal(Resp, &plans)
	if err != nil {
		log.Fatal(err)
	}
	for i := range plans {
		if plans[i].Start_date.Month() == cpr.Start_date.Month() && plans[i].Start_date.Day() == cpr.Start_date.Day() {
			if plans[i].Start_date.Hour() == cpr.Start_date.Hour() {
				fmt.Println("You have already plan in This hour")
				return false
			}
			if int(plans[i].Start_date.Hour()) < int(cpr.Start_date.Hour()) && int(plans[i].Finish_date.Hour()) > int(cpr.Start_date.Hour()) {
				fmt.Println("You have already plan in This hour")
				return false
			}
		}

	}
	return true
}
func GetResponseResult(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(respData)
	return respData
}
