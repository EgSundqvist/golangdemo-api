package controllers

import (
	"net/http"
	"strconv"

	"github.com/EgSundqvist/data"
	"github.com/EgSundqvist/models"
	"github.com/gin-gonic/gin"
)

func HandleGetAllEmployees(c *gin.Context) {
	employees := data.GetAllEmployees()
	c.JSON(http.StatusOK, employees)
}

func HandleGetEmployeeById(c *gin.Context) {
	id := c.Param("id")
	employeeId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	employee := data.GetEmployeeById(employeeId)

	if employee == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	c.JSON(http.StatusOK, employee)
}

func HandleCreateEmployee(c *gin.Context) {
	var newEmployee models.Employee
	if err := c.BindJSON(&newEmployee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	data.CreateEmployee(newEmployee)
	c.JSON(http.StatusCreated, newEmployee)
}

func HandleUpdateEmployee(c *gin.Context) {
	id := c.Param("id")                 // Extract the ID from the URL parameter
	employeeId, err := strconv.Atoi(id) // Convert the ID to an integer
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var employee models.Employee
	if err := c.BindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Ensure the ID in the URL matches the ID in the JSON body
	if employee.Id != employeeId {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID mismatch"})
		return
	}

	// Check if the employee exists
	existingEmployee := data.GetEmployeeById(employeeId)
	if existingEmployee == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	if err := data.UpdateEmployee(employee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update employee"})
		return
	}

	c.JSON(http.StatusOK, employee)
}

func HandleDeleteEmployee(c *gin.Context) {
	id := c.Param("id")            // Extract the ID from the URL parameter
	numId, err := strconv.Atoi(id) // Convert the ID to an integer
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	employeeFromDb := data.GetEmployeeById(numId)
	if employeeFromDb == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	data.DeleteEmployee(employeeFromDb)
	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
}
