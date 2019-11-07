package main

import (
	"github.com/jqhnmaina/Product-Catalogue-API/config"
	"github.com/jqhnmaina/Product-Catalogue-API/database/migrations"
	"github.com/jqhnmaina/Product-Catalogue-API/handlers"
)

func main() {
	config.InitConfig()
	migrations.Migrate()

	handlers.Init()
}
