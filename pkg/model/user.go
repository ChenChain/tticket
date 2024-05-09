package model

import "time"

type User struct {
	ID           int64 `gorm:"primary_key" json:"id"`
	Name         string
	Mail         string
	Created_Time time.Time
	Updated_Time time.Time
}
