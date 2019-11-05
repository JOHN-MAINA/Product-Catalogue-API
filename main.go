package main

import (
	"github.com/jqhnmaina/Product-Catalogue-API/database/migrations"
	"github.com/jqhnmaina/Product-Catalogue-API/handlers"
)

func main() {
	migrations.Migrate()

	handlers.Init()
}
