package database

import (
	"context"
	_ "database/sql"
	"fmt"

	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	"truenorth/packages/logger"
	"truenorth/packages/utils"
)

func RunMigrations(ctx context.Context) {
	dbInstance := GetInstance()
	dbConn, err := dbInstance.DB()
	if err != nil {
		panic(err)
	}
	driver, err := postgres.WithInstance(dbConn, &postgres.Config{})
	if err != nil {
		panic(err)
	}
	migrationPath := utils.GetEnv("MIGRATION_PATH", "")
	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://./../sql/migrations/%s", migrationPath),
		"postgres", driver)

	if err != nil {
		logger.GetLogger().Errorw(ctx, "Error on migrations", "value", err.Error())
		panic(err)
	}
	err = m.Up()
	if err != nil && err.Error() != "no change" {
		logger.GetLogger().Errorw(ctx, "Error on migrations", "value", err.Error())
		panic(err)
	}
	logger.GetLogger().Infow(ctx, "Migrations run successfully")
}
