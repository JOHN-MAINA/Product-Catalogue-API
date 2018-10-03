package main

import (
	"./database/migrations"
	"./handlers"
)

func main() {
	migrations.Migrate()

	handlers.Init()
}
