package data

import (
	"github.com/EgSundqvist/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(file, server, database, username, password string, port int) {
	var err error
	db, err = gorm.Open(sqlite.Open(file), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.Employee{}, &models.ApiAdmin{})

	var numEmp int64
	db.Model(&models.Employee{}).Count(&numEmp)
	if numEmp == 0 {
		employees := []models.Employee{
			{Id: 1, Age: 20, City: "Stockholm", Name: "Jerry"},
			{Id: 2, Age: 30, City: "Stockholm", Name: "Conny"},
			{Id: 3, Age: 40, City: "Stockholm", Name: "Ronny"},
		}

		for _, emp := range employees {
			db.Create(&emp)
		}
	}

	var numApiAdmin int64
	db.Model(&models.ApiAdmin{}).Count(&numApiAdmin)
	if numApiAdmin == 0 {
		// Create an initial ApiAdmin with the first employee
		_, err := CreateApiAdmin(1)
		if err != nil {
			panic("failed to create initial Apiadmin")
		}
	}
}
