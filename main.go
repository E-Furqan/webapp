package main

import (
	"github.com/E-Furqan/webapp/controllers"
	"github.com/E-Furqan/webapp/database"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connection()
	server := gin.Default()
	server.GET("/student", controllers.Getdata)
	server.POST("/student", controllers.Create_student)
	server.DELETE("/delstudent/:id", controllers.Delete_student)
	server.PATCH("/updateinfo/:id", controllers.Update_info)

	server.Run(":8081")
}
