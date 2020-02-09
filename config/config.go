package config

import (
	"log"
	"os"
	"strconv"
)

var DbHost = ""

var DbPort = ""

var DbName = ""

var DbUser = ""

var DbPass = ""

// postgres mysql sqlite
var DbDialect = ""

var AppPort = 3001

func InitConfig() {
	if host := os.Getenv("DB_HOST"); len(host) == 0 {
		log.Println("Init config variables: env var DB_HOST is not set")
	} else {
		DbHost = host
	}

	if port := os.Getenv("DB_PORT"); len(port) == 0 {
		log.Println("Init config variables: env var DB_PORT is not set")
	} else {
		DbPort = port
	}

	if dbName := os.Getenv("DB_NAME"); len(dbName) == 0 {
		log.Println("Init config variables: env var DB_NAME is not set")
	} else {
		DbName = dbName
	}

	if dbUser := os.Getenv("DB_USER"); len(dbUser) == 0 {
		log.Println("Init config variables: env var DB_USER is not set")
	} else {
		DbUser = dbUser
	}

	if dbPass := os.Getenv("DB_PASSWORD"); len(dbPass) == 0 {
		log.Println("Init config variables: env var DB_PASSWORD is not set")
	} else {
		DbPass = dbPass
	}

	if dbDialect := os.Getenv("DB_DIALECT"); len(dbDialect) == 0 {
		log.Println("Init config variables: env var DB_DIALECT is not set")
	} else {
		DbDialect = dbDialect
	}

	if port := os.Getenv("PORT"); len(port) == 0 {
		log.Println("Init config variables: env var PORT is not set using 3001 as default port")
	} else {
		AppPort, _ = strconv.Atoi(port)
	}
}
