package data

import (
	"fmt"
	"time"

	"github.com/EgSundqvist/models"
)

func CreateApiAdmin(employeeId int) (*models.ApiAdmin, error) {
	employee := GetEmployeeById(employeeId)
	if employee == nil {
		return nil, fmt.Errorf("employee not found")
	}
	apiadmin := &models.ApiAdmin{
		EmployeeID: uint(employee.Id),
		SuperUser:  true,
		AdminFrom:  time.Now(),
		AdminTo:    time.Now().AddDate(1, 0, 0),
		ExtraInfo:  "New API Admin",
	}
	result := db.Create(apiadmin)
	return apiadmin, result.Error
}

func GetApiAdmin() *models.ApiAdmin {
	var apiadmin models.ApiAdmin
	db.Preload("Employee").Order("id desc").First(&apiadmin)
	if apiadmin.ID == 0 {
		return nil
	}
	return &apiadmin
}
