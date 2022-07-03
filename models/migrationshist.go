package models

type Migrations struct {
	Migrated string `gorm:"type:varchar(10)"`
}
