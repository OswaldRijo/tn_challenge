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

var StatusIntToNameMap = map[StatusEnum]string{
	StatusActive:   "ACTIVE",
	StatusInactive: "INACTIVE",
}
var statusNameToValueMap = map[string]StatusEnum{
	"ACTIVE":   StatusActive,
	"INACTIVE": StatusInactive,
}

func (u StatusEnum) Scan(value interface{}) error {
	u = statusNameToValueMap[value.(string)]
	return nil
}

func (u StatusEnum) Value() (driver.Value, error) {
	statusName := StatusIntToNameMap[u]
	return statusName, nil
}

type User struct {
	ID        int64  `gorm:"primaryKey"`
	Username  string `gorm:"index"`
	Password  string
	Status    StatusEnum
	CreatedAt time.Time
	UpdatedAt time.Time
}
