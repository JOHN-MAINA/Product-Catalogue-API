package migrations

import "github.com/jqhnmaina/Product-Catalogue-API/database/database"

func Migrate() {

	db := database.Connect()

	db.AutoMigrate(&Product{}, &Category{})

}

func DropTables() {

	db := database.Connect()

	db.DropTableIfExists(&Product{}, &Category{})

	database.CloseConnection(db)
}
