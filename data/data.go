package data

import (
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

	CreateApiAdmin()
}

func CreateApiAdmin() {
	employee := GetEmployeeById(1)
	if employee == nil {
		return
	}
	me := &Apiadmin{
		EmployeeID: uint(employee.Id),
		SuperUser:  true,
		AdminFrom:  time.Now(),
		AdminTo:    time.Now().AddDate(1, 0, 0),
		ExtraInfo:  "Kingarnas king på att administera API:er",
	}
	db.Create(me)
}

func GetApiAdmin() *Apiadmin {
	var apiadmin Apiadmin
	db.Preload("Employee").First(&apiadmin, "employee_id = ?", 1)
	if apiadmin.ID == 0 {
		return nil
	}
	return &apiadmin
}
