package database

import (
	"fmt"
	"github.com/jqhnmaina/Product-Catalogue-API/config"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB gorm.DB

func Connect() gorm.DB {

	defer recover()
	dbURL := ""

	switch config.DbDialect {
	case "postgres":
		dbURL = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", config.DbHost, config.DbPort, config.DbUser, config.DbName, config.DbPass)
		break
	case "mysql":
		dbURL = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", config.DbUser, config.DbPass, config.DbName)
		break
	case "sqlite":
		dbURL = config.DbName
		break
	}
	DB, err := gorm.Open(config.DbDialect, dbURL)

	if err != nil {
		log.Fatalln("Failed to connect to db: " + err.Error())
	}

	return *DB
}

func CloseConnection(db gorm.DB) {
	db.Close()
}
