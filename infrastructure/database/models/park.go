package models

import "gorm.io/gorm"

type ParkVague struct {
	gorm.Model
	Name  string
	Vague int32
	Limit int32
}
