package config

import (
	"truenorth/packages/utils"
)

func loadDev() {
	Config.Port = utils.GetEnv("PORT")
	Config.DBHost = utils.GetEnv("DB_HOST", "localhost")
	Config.DBPort = utils.GetEnv("DB_PORT", "5432")
	Config.DBUser = utils.GetEnv("DB_USER", "user")
	Config.DBPass = utils.GetEnv("DB_PASS", "pass")
	Config.DBName = utils.GetEnv("DB_NAME", "truenorth-db")
	Config.ENV = "dev"
	Config.AppName = utils.GetEnv("APP_NAME", "operations_service")
	debugMode := utils.GetEnv("DEBUG", "false")
	Config.Debug = debugMode == "true"
	runMigrations := utils.GetEnv("RUN_MIGRATIONS", "false")
	Config.RunMigrations = runMigrations == "true"
}
