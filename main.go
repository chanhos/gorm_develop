package main

import (
	"fmt"
	"log"
	"time"

	"github.com/chanhos/workple_demo/db"
	"github.com/chanhos/workple_demo/models"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := db.ConnDB("dev")

	if err != nil {
		log.Fatal(err.Error())
	}

	ExcuteMigrations(db)

	//defModel := gorm.Model{CreatedAt: time.Now()}

	//card1 := models.CreditCard{Model: gorm.Model{CreatedAt: time.Now()}, Number: "0000-0000-0000", UserName: "정찬호"}

	//user1 := models.User{Model: gorm.Model{CreatedAt: time.Now()}, CreditCard: card1, Name: "정찬호"}

	wpl1 := models.WorkPlaces{WorkplaceName: "채소가게", WorkplacePhoneNo: "031-666-7777", CorporateNumber: "0112-123123-10", AccountsUserID: "chjung", BaseModel: models.BaseModel{CreatedAt: time.Now()}}
	wpl2 := models.WorkPlaces{WorkplaceName: "과일가게", WorkplacePhoneNo: "031-666-7777", CorporateNumber: "0112-123123-10", AccountsUserID: "chjung", BaseModel: models.BaseModel{CreatedAt: time.Now()}}

	workplaces := []models.WorkPlaces{}
	workplaces = append(workplaces, wpl1, wpl2)

	acc1 := models.Accounts{UserName: "정찬호", UserPassword: "110356jm!", MobileNo: "010-4705-8716", UserID: "chjung",
		BaseModel: models.BaseModel{CreatedAt: time.Now()}, WorkPlaces: workplaces}

	rslt := db.Create(&acc1)

	fmt.Println(rslt.RowsAffected)
	my := &models.Accounts{}
	db.Preload("WorkPlaces").Where("user_id = ?", "chjung").Find(my)

	fmt.Println(len(my.WorkPlaces))
}
