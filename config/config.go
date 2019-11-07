package config

import (
	"os"
	"strconv"
)

var DbHost = os.Getenv("DB_HOST")

var DbPort, _ = strconv.Atoi(os.Getenv("DB_PORT"))

var DbName = os.Getenv("DB_NAME")

var DbUser = os.Getenv("DB_USER")

var DbPass = os.Getenv("DB_PASSWORD")

// postgres mysql
var DbDialect = os.Getenv("DB_DIALECT")

var AppPort = 3001
