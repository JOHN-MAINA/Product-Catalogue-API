package database

import (
	"fmt"
	"github.com/JOHN-MAINA/Product-Catalogue-API/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB gorm.DB

func Connect() gorm.DB {
	defer recover()
	dbURL := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", config.DbUser, config.DbPass, config.DbName)
	db, err := gorm.Open("mysql", dbURL)

	if err != nil {
		panic(err.Error())
	}

	//defer db.Close()
	return *db

}

func CloseConnection(db gorm.DB) {
	db.Close()
}
