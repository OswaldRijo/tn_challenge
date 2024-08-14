package config

import (
	"truenorth/packages/utils"
)

type configuration struct {
	Port                        string
	DBHost                      string
	DBPort                      string
	DBUser                      string
	DBPass                      string
	DBName                      string
	ENV                         string
	AppName                     string
	Debug                       bool
	RunMigrations               bool
	DefaultUserBalance          float64
	AdditionOperationCost       float64
	SubtractionOperationCost    float64
	DivisionOperationCost       float64
	MultiplicationOperationCost float64
	RandomStringOperationCost   float64
	SquareRootOperationCost     float64
	UserCreatedQueue            string
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
