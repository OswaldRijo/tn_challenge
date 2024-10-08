package config

import (
	"truenorth/packages/utils"
)

func loadProd() {
	Config.UserCreatedTopicArn = utils.GetEnv("USER_CREATED_TOPIC_ARN", "")
	Config.Salt = utils.GetEnv("HASH_SALT")
	Config.Port = utils.GetEnv("PORT")
	Config.DBHost = utils.GetEnv("DB_HOST")
	Config.DBPort = utils.GetEnv("DB_PORT")
	Config.DBUser = utils.GetEnv("DB_USER")
	Config.DBPass = utils.GetEnv("DB_PASS")
	Config.DBName = utils.GetEnv("DB_NAME")
	Config.ENV = "prod"
	Config.AppName = utils.GetEnv("APP_NAME", "users_service")
	Config.Debug = false
	runMigrations := utils.GetEnv("RUN_MIGRATIONS", "false")
	Config.RunMigrations = runMigrations == "true"
}
