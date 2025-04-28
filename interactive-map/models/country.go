package models

import "gorm.io/gorm"

type CountryStatus struct {
	gorm.Model
	CountryISO string
	Status     string // "visited", "want"
	UserID     uint
}
