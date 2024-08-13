package models

import (
	"time"
)

type Record struct {
	ID                int64 `gorm:"primaryKey"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	OperationID       int64 `gorm:"index"`
	UserID            int64 `gorm:"index"`
	Amount            float64
	UserBalance       float64
	Deleted           bool
	OperationResponse []byte

	Operation *Operation `gorm:"references:OperationID"`
}

func (r *Record) SetCreatedAt(createdAt time.Time) *Record {
	r.CreatedAt = createdAt
	return r
}

func (r *Record) SetUpdatedAt(updatedAt time.Time) *Record {
	r.UpdatedAt = updatedAt
	return r
}

func (r *Record) SetOperationID(operationID int64) *Record {
	r.OperationID = operationID
	return r
}

func (r *Record) SetUserID(userID int64) *Record {
	r.UserID = userID
	return r
}

func (r *Record) SetAmount(amount float64) *Record {
	r.Amount = amount
	return r
}

func (r *Record) SetUserBalance(userBalance float64) *Record {
	r.UserBalance = userBalance
	return r
}

func (r *Record) SetDeleted(deleted bool) *Record {
	r.Deleted = deleted
	return r
}

func (r *Record) SetOperationResponse(operationResponse []byte) *Record {
	r.OperationResponse = operationResponse
	return r
}

func NewRecord() *Record {
	return &Record{}
}
