package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

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
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	db.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	db.SetMaxOpenConns(20)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	db.SetConnMaxLifetime(2 * time.Hour)

	return nil
}

func getDbConn(url string) *sql.DB {
	dbConn, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal(err)
	}

	return dbConn
}

func GetInstance() *sql.DB {
	if db == nil {
		err := Init()
		if err != nil {
			log.Fatal("cannot connect to db: ", err)
		}
		return db
	}
	return db
}

var db *sql.DB
