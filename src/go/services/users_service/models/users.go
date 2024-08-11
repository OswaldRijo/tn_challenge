package models

import (
	"database/sql/driver"
	"time"
)

type StatusEnum int32

const (
	StatusActive StatusEnum = iota
	StatusInactive
)

func (u *StatusEnum) Scan(value interface{}) error {
	*u = StatusEnum(value.(int32))
	return nil
}

func (u *StatusEnum) Value() (driver.Value, error) {
	return int32(*u), nil
}

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"index"`
	Password  string
	Status    StatusEnum
	CreatedAt time.Time
	UpdatedAt time.Time
}
