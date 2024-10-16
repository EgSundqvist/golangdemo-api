package data

import (
	"fmt"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var employees []Employee

func GetAllEmployees() []Employee {
	var employees []Employee
	db.Find(&employees)
	return employees
}
func GetEmployeeById(id int) *Employee {
	var employee Employee
	db.First(&employee, id)
	if employee.Id == 0 {
		return nil
	}
	return &employee
}

func CreateEmployee(newEmployee Employee) {
	db.Create(&newEmployee)
}

func UpdateEmployee(employee Employee) error {
	result := db.Save(&employee)
	return result.Error
}

func DeleteEmployee(employee *Employee) {
	db.Delete(employee)
}

var db *gorm.DB

func Init(file, server, database, username, password string, port int) {
	var err error
	db, err = gorm.Open(sqlite.Open(file), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Employee{}, &Apiadmin{})

	var numEmp int64
	db.Model(&Employee{}).Count(&numEmp)
	if numEmp == 0 {
		employees = append(employees, Employee{Id: 1, Age: 20, City: "Stockholm", Name: "Jerry"})
		employees = append(employees, Employee{Id: 2, Age: 30, City: "Stockholm", Name: "Conny"})
		employees = append(employees, Employee{Id: 3, Age: 40, City: "Stockholm", Name: "Ronny"})

		for _, emp := range employees {
			db.Create(&emp)
		}
	}

	var numApiAdmin int64
	db.Model(&Apiadmin{}).Count(&numApiAdmin)
	if numApiAdmin == 0 {
		// Create an initial Apiadmin with the first employee
		_, err := CreateApiAdmin(1)
		if err != nil {
			panic("failed to create initial Apiadmin")
		}
	}
}

func CreateApiAdmin(employeeId int) (*Apiadmin, error) {
	employee := GetEmployeeById(employeeId)
	if employee == nil {
		return nil, fmt.Errorf("Employee not found")
	}
	apiadmin := &Apiadmin{
		EmployeeID: uint(employee.Id),
		SuperUser:  true,
		AdminFrom:  time.Now(),
		AdminTo:    time.Now().AddDate(1, 0, 0),
		ExtraInfo:  "New API Admin",
	}
	result := db.Create(apiadmin)
	return apiadmin, result.Error
}

func GetApiAdmin() *Apiadmin {
	var apiadmin Apiadmin
	db.Preload("Employee").Order("id desc").First(&apiadmin)
	if apiadmin.ID == 0 {
		return nil
	}
	return &apiadmin
}
