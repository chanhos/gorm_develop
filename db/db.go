package db

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnDB(stage string) (*gorm.DB, error) {

	dsn, err := loadConnStrByEnv(stage)

	if err != nil {
		return nil, errors.New("connection error")
	}

	config := &gorm.Config{}
	config.Logger = logger.Default.LogMode(logger.Info)

	conn, err := gorm.Open(mysql.Open(dsn), config)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	return conn, nil
}

func loadConnStrByEnv(stage string) (connstr string, err error) {
	connstr = ""
	switch stage {
	case "dev":
		err = godotenv.Load("dev.env")

		if err != nil {
			return
		}

		connstr = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_SCHEMA"))

		return
	default:
		return "", nil
	}
}
