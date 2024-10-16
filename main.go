package main

import (
	"net/http"
	"strconv"

	"github.com/EgSundqvist/data"
	"github.com/gin-gonic/gin"
)

func handleStart(c *gin.Context) {
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("Hello, World!"))
}

func handleAbout(c *gin.Context) {
	c.HTML(http.StatusOK, "about.html", gin.H{})
}

func handleGetAllEmployees(c *gin.Context) {
	employees := data.GetAllEmployees()
	c.JSON(http.StatusOK, employees)
}

func handleGetEmployeeById(c *gin.Context) {
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

func handleCreateEmployee(c *gin.Context) {
	var newEmployee data.Employee
	if err := c.BindJSON(&newEmployee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	data.CreateEmployee(newEmployee)
	c.JSON(http.StatusCreated, newEmployee)
}

func handleUpdateEmployee(c *gin.Context) {
	id := c.Param("id")                 // Extract the ID from the URL parameter
	employeeId, err := strconv.Atoi(id) // Convert the ID to an integer
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var employee data.Employee
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

func handleDeleteEmployee(c *gin.Context) {
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

func handleGetApiAdmin(c *gin.Context) {
	me := data.GetApiAdmin()
	if me == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Me not found"})
		return
	}
	c.JSON(http.StatusOK, me)
}

var config Config

func main() {

	readConfig(&config)
	data.Init(config.Database.File,
		config.Database.Server,
		config.Database.Database,
		config.Database.Username,
		config.Database.Password,
		config.Database.Port)

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", handleStart)
	r.GET("/api/about", handleAbout)
	r.GET("/api/employee", handleGetAllEmployees)
	r.GET("/api/employee/:id", handleGetEmployeeById)
	r.POST("/api/employee", handleCreateEmployee)
	r.PUT("/api/employee/:id", handleUpdateEmployee)
	r.DELETE("/api/employee/:id", handleDeleteEmployee)
	r.GET("/api/apiadmin", handleGetApiAdmin)

	r.Run()
}
