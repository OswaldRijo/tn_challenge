package config

import (
	"strconv"

	"truenorth/packages/utils"
)

func loadLocal() {
	defaultUserBalanceStr := utils.GetEnv("DEFAULT_USER_BALANCE", "500")
	defaultUserBalance, err := strconv.ParseFloat(defaultUserBalanceStr, 32)
	if defaultUserBalance == 0 || err != nil {
		panic("Wrong value for DEFAULT_USER_BALANCE: " + defaultUserBalanceStr)
	}

	additionOperationCostStr := utils.GetEnv("ADDITION_OPERATION_COST", "1")
	additionOperationCost, err := strconv.ParseFloat(additionOperationCostStr, 32)
	if additionOperationCost == 0 || err != nil {
		panic("Wrong value for ADDITION_OPERATION_COST: " + additionOperationCostStr)
	}

	subtractionOperationCostStr := utils.GetEnv("SUBTRACTION_OPERATION_COST", "1")
	subtractionOperationCost, err := strconv.ParseFloat(subtractionOperationCostStr, 32)
	if subtractionOperationCost == 0 || err != nil {
		panic("Wrong value for SUBTRACTION_OPERATION_COST: " + subtractionOperationCostStr)
	}

	divisionOperationCostStr := utils.GetEnv("DIVISION_OPERATION_COST", "3")
	divisionOperationCost, err := strconv.ParseFloat(divisionOperationCostStr, 32)
	if divisionOperationCost == 0 || err != nil {
		panic("Wrong value for DIVISION_OPERATION_COST: " + divisionOperationCostStr)
	}

	squareRootOperationCostStr := utils.GetEnv("SQUARE_ROOT_OPERATION_COST", "5")
	squareRootOperationCost, err := strconv.ParseFloat(squareRootOperationCostStr, 32)
	if squareRootOperationCost == 0 || err != nil {
		panic("Wrong value for SQUARE_ROOT_OPERATION_COST: " + squareRootOperationCostStr)
	}

	multiplicationOperationCostStr := utils.GetEnv("MULTIPLICATION_OPERATION_COST", "3")
	multiplicationOperationCost, err := strconv.ParseFloat(multiplicationOperationCostStr, 32)
	if multiplicationOperationCost == 0 || err != nil {
		panic("Wrong value for MULTIPLICATION_OPERATION_COST: " + multiplicationOperationCostStr)
	}

	randomStringOperationCostStr := utils.GetEnv("RANDOM_STRING_OPERATION_COST", "5")
	randomStringOperationCost, err := strconv.ParseFloat(randomStringOperationCostStr, 32)
	if randomStringOperationCost == 0 || err != nil {
		panic("Wrong value for RANDOM_STRING_OPERATION_COST: " + randomStringOperationCostStr)
	}

	Config.Port = utils.GetEnv("PORT", "11001")
	Config.DBHost = utils.GetEnv("DB_HOST", "localhost")
	Config.DBPort = utils.GetEnv("DB_PORT", "5432")
	Config.DBUser = utils.GetEnv("DB_USER", "user")
	Config.DBPass = utils.GetEnv("DB_PASS", "pass")
	Config.DBName = utils.GetEnv("DB_NAME", "truenorth_db")
	Config.ENV = "local"
	Config.AppName = utils.GetEnv("APP_NAME", "operations_service")
	debugMode := utils.GetEnv("DEBUG", "true")
	Config.Debug = debugMode == "true"
	runMigrations := utils.GetEnv("RUN_MIGRATIONS", "false")
	Config.UserCreatedQueue = utils.GetEnv("USER_CREATED_QUEUE", "user-created-consumer-dev")
	Config.RunMigrations = runMigrations == "true"
	Config.DefaultUserBalance = defaultUserBalance
	Config.AdditionOperationCost = additionOperationCost
	Config.DivisionOperationCost = divisionOperationCost
	Config.MultiplicationOperationCost = multiplicationOperationCost
	Config.RandomStringOperationCost = randomStringOperationCost
	Config.SquareRootOperationCost = squareRootOperationCost
	Config.SubtractionOperationCost = subtractionOperationCost
}
