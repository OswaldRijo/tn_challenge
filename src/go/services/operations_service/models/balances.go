package models

import "time"

type Balance struct {
	ID             int64 `gorm:"primaryKey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	UserID         int64 `gorm:"index"`
	CurrentBalance float64
}

func (op *Balance) SetCreatedAt(createdAt time.Time) *Balance {
	op.CreatedAt = createdAt
	return op
}

func (op *Balance) SetUpdatedAt(updatedAt time.Time) *Balance {
	op.UpdatedAt = updatedAt
	return op
}

func (op *Balance) SetUserID(userID int64) *Balance {
	op.UserID = userID
	return op
}

func (op *Balance) SetCurrentBalance(currentBalance float64) *Balance {
	op.CurrentBalance = currentBalance
	return op
}

func NewBalance() *Balance {
	return &Balance{}
}
