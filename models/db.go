package models

import (
	"fmt"
	"github.com/ChallenAi/iTools-service/conf"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func InitDB() *gorm.DB {

	db, err := gorm.Open("postgres",
		"sslmode=disable"+
			" host="+conf.Db_host+
			" port="+conf.Db_port+
			" user="+conf.Db_username+
			" dbname="+conf.Db_name+
			" password="+conf.Db_password)
	if err != nil {
		fmt.Println("db err: ", err)
	}
	db.DB().SetMaxIdleConns(10)
	db.LogMode(false)

	DB = db
	return DB
}
