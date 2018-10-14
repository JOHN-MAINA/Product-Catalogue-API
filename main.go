package main

import (
	"github.com/JOHN-MAINA/Product-Catalogue-API/database/migrations"
	"github.com/JOHN-MAINA/Product-Catalogue-API/handlers"
)

func main() {
	migrations.Migrate()

	handlers.Init()
}
