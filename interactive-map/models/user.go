package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string
	Email     string `gorm:"uniqueIndex"`
	Password  string
	Countries []CountryStatus
}
