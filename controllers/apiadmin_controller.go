package controllers

import (
	"net/http"

	"github.com/EgSundqvist/data"
	"github.com/gin-gonic/gin"
)

func HandleGetApiAdmin(c *gin.Context) {
	me := data.GetApiAdmin()
	if me == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Me not found"})
		return
	}
	c.JSON(http.StatusOK, me)
}

func HandleCreateApiAdmin(c *gin.Context) {
	var request struct {
		EmployeeID int `json:"employeeId"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	apiadmin, err := data.CreateApiAdmin(request.EmployeeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, apiadmin)
}
