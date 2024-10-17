package routes

import (
	"net/http"

	"github.com/EgSundqvist/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("Hello, World!"))
	})
	r.GET("/api/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about.html", gin.H{})
	})

	r.GET("/api/employee", controllers.HandleGetAllEmployees)
	r.GET("/api/employee/:id", controllers.HandleGetEmployeeById)
	r.POST("/api/employee", controllers.HandleCreateEmployee)
	r.PUT("/api/employee/:id", controllers.HandleUpdateEmployee)
	r.DELETE("/api/employee/:id", controllers.HandleDeleteEmployee)

	r.GET("/api/apiadmin", controllers.HandleGetApiAdmin)
	r.POST("/api/apiadmin", controllers.HandleCreateApiAdmin)

	return r
}
