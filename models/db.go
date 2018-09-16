package models

import (
	"fmt"
	"github.com/ChallenAi/iTools-service/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func InitDB() *gorm.DB {

	db, err := gorm.Open("postgres",
		"sslmode=disable"+
			" host="+config.DB_host+
			" port="+config.DB_port+
			" user="+config.DB_username+
			" dbname="+config.DBname+
			" password="+config.DB_password)
	if err != nil {
		fmt.Println("db err: ", err)
	}
	db.DB().SetMaxIdleConns(10)
	db.LogMode(true)

	DB = db
	return DB
}
