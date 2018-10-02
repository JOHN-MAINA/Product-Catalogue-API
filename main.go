package main

import "./database/migrations"

func main() {
	migrations.Migrate()
}
