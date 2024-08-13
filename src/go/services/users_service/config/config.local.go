package config

import (
	"truenorth/packages/utils"
)

func loadLocal() {
	Config.UserCreatedTopicArn = utils.GetEnv("USER_CREATED_TOPIC_ARN", "")
	Config.Salt = utils.GetEnv("HASH_SALT", "3354961358384ce3a1436527f2825844")
	Config.Port = utils.GetEnv("PORT", "11001")
	Config.DBHost = utils.GetEnv("DB_HOST", "localhost")
	Config.DBPort = utils.GetEnv("DB_PORT", "5432")
	Config.DBUser = utils.GetEnv("DB_USER", "user")
	Config.DBPass = utils.GetEnv("DB_PASS", "pass")
	Config.DBName = utils.GetEnv("DB_NAME", "truenorth-db")
	Config.ENV = "local"
	Config.AppName = utils.GetEnv("APP_NAME", "users_service")
	debugMode := utils.GetEnv("DEBUG", "true")
	Config.Debug = debugMode == "true"
	runMigrations := utils.GetEnv("RUN_MIGRATIONS", "false")
	Config.RunMigrations = runMigrations == "true"

}
