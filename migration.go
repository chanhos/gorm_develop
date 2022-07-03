package main

import (
	"github.com/chanhos/workple_demo/models"
	"gorm.io/gorm"
)

func ExcuteMigrations(db *gorm.DB) {
	mig := db.Migrator()

	if tables, err := mig.GetTables(); err == nil {
		find := false
		for _, table := range tables {
			if table == "migrations" {
				find = true
			}
		}
		if !find {
			mig.CreateTable(&models.Migrations{})
		}
	}

	mig.DropTable(&models.User{})

	mig.DropConstraint(&models.CreditCard{}, "fk_users_credit_card")
	mig.DropTable(&models.CreditCard{})

	mig.DropTable(&models.Accounts{})

	mig.DropConstraint(&models.WorkPlaces{}, "fk_accounts_work_places")
	mig.DropTable(&models.WorkPlaces{})

	db.AutoMigrate(&models.User{}, &models.CreditCard{}, &models.Accounts{}, &models.WorkPlaces{})

}
