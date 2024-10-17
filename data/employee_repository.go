package data

import (
	"github.com/EgSundqvist/models"
)

func GetAllEmployees() []models.Employee {
	var employees []models.Employee
	db.Find(&employees)
	return employees
}

func GetEmployeeById(id int) *models.Employee {
	var employee models.Employee
	db.First(&employee, id)
	if employee.Id == 0 {
		return nil
	}
	return &employee
}

func CreateEmployee(newEmployee models.Employee) {
	db.Create(&newEmployee)
}

func UpdateEmployee(employee models.Employee) error {
	result := db.Save(&employee)
	return result.Error
}

func DeleteEmployee(employee *models.Employee) {
	db.Delete(employee)
}
