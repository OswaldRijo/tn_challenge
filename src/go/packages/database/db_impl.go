package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"truenorth/packages/utils"
)

func Init() error {
	dbHost := utils.GetEnv("DB_HOST")
	dbPort := utils.GetEnv("DB_PORT")
	dbUser := utils.GetEnv("DB_USER")
	dbPass := utils.GetEnv("DB_PASS")
	dbName := utils.GetEnv("DB_NAME")
	url := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName)

	db = getDbConn(url)

	return nil
}

func getDbConn(url string) *gorm.DB {
	dbConn, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return dbConn
}

func GetInstance() *gorm.DB {
	if db == nil {
		err := Init()
		if err != nil {
			log.Fatal("cannot connect to db: ", err)
		}
		return db
	}
	return db
}

var db *gorm.DB
