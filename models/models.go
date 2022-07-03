package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name       string     `gorm:"type:varchar(15);not null;unique;index"`
	CreditCard CreditCard `gorm:"foreignkey:UserName;references:name"`
}

type CreditCard struct {
	gorm.Model
	Number   string `gorm:"type:varchar(20);not null;" json:"number"`
	UserName string `gorm:"type:varchar(15);not null;unique"`
}

type BaseModel struct {
	CreatedAt time.Time      `gorm:"type:datetime;"`
	UpdatedAt time.Time      `gorm:"type:datetime;"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Accounts struct {
	AccID        uint         `gorm:"primarykey"`
	UserPassword string       `gorm:"type:varchar(128)"`
	UserName     string       `gorm:"type:varchar(30)"`
	MobileNo     string       `gorm:"type:varchar(15)"`
	UserID       string       `gorm:"type:varchar(15);not null;unique;index"`
	WorkPlaces   []WorkPlaces `gorm:"foreignkey:AccountsUserID;references:user_id"`
	BaseModel
}

type WorkPlaces struct {
	WplID            uint   `gorm:"primarykey"`
	WorkplaceName    string `gorm:"type:varchar(200)"`
	WorkplacePhoneNo string `gorm:"type:varchar(15)"`
	CorporateNumber  string `gorm:"type:varchar(20)"`
	AccountsUserID   string `gorm:"type:varchar(15);column:owner_id;not null"`
	BaseModel
}
