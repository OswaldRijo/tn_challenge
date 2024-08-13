package models

import (
	"database/sql/driver"
	"time"

	operationspb "truenorth/pb/operations"
)

type OperationTypeEnum int32

const (
	OperationTypeAddition OperationTypeEnum = iota
	OperationTypeSubtraction
	OperationTypeMultiplication
	OperationTypeDivision
	OperationTypeSquareRoot
	OperationTypeRandomString
)

var statusIntToNameMap = map[OperationTypeEnum]string{
	OperationTypeAddition:       "ADDITION",
	OperationTypeSubtraction:    "SUBTRACTION",
	OperationTypeMultiplication: "MULTIPLICATION",
	OperationTypeDivision:       "DIVISION",
	OperationTypeSquareRoot:     "SQUARE_ROOT",
	OperationTypeRandomString:   "RANDOM_STRING",
}
var statusNameToValueMap = map[string]OperationTypeEnum{
	"ADDITION":       OperationTypeAddition,
	"SUBTRACTION":    OperationTypeSubtraction,
	"MULTIPLICATION": OperationTypeMultiplication,
	"DIVISION":       OperationTypeDivision,
	"SQUARE_ROOT":    OperationTypeSquareRoot,
	"RANDOM_STRING":  OperationTypeRandomString,
}

func (u OperationTypeEnum) Scan(value interface{}) error {
	u = statusNameToValueMap[value.(string)]
	return nil
}

func (u OperationTypeEnum) Value() (driver.Value, error) {
	statusName := statusIntToNameMap[u]
	return statusName, nil
}

func (u OperationTypeEnum) ToPb() operationspb.OperationType {
	return operationspb.OperationType(u)
}

func OpTypeFromPb(operationTypePb operationspb.OperationType) OperationTypeEnum {
	return OperationTypeEnum(operationTypePb)
}

type Operation struct {
	ID            int64 `gorm:"primaryKey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	UserID        int64 `gorm:"index"`
	Cost          float64
	OperationType OperationTypeEnum
	Args          []byte
}

func (op *Operation) SetCreatedAt(createdAt time.Time) *Operation {
	op.CreatedAt = createdAt
	return op
}

func (op *Operation) SetUpdatedAt(updatedAt time.Time) *Operation {
	op.UpdatedAt = updatedAt
	return op
}

func (op *Operation) SetUserID(userID int64) *Operation {
	op.UserID = userID
	return op
}

func (op *Operation) SetCost(cost float64) *Operation {
	op.Cost = cost
	return op
}

func (op *Operation) SetOperationType(operationType OperationTypeEnum) *Operation {
	op.OperationType = operationType
	return op
}

func (op *Operation) SetArgs(args []byte) *Operation {
	op.Args = args
	return op
}

func NewOperation() *Operation {
	return &Operation{}
}
