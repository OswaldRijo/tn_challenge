package config

import (
	"truenorth/packages/utils"
)

type configuration struct {
	Port          string
	DBHost        string
	DBPort        string
	DBUser        string
	DBPass        string
	DBName        string
	Salt          string
	ENV           string
	AppName       string
	Debug         bool
	RunMigrations bool
}

var Config = configuration{}

func Load() (err error) {

	switch env := utils.GetEnv("ENV"); env {
	case "prod":
		loadProd()
	case "dev":
		loadDev()
	default:
		loadLocal()
	}
	return
}
