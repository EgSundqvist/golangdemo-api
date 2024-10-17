package models

import "time"

type ApiAdmin struct {
	ID         uint `gorm:"primaryKey"`
	EmployeeID uint
	Employee   Employee `gorm:"foreignKey:EmployeeID"`
	SuperUser  bool
	AdminFrom  time.Time
	AdminTo    time.Time
	ExtraInfo  string
}
