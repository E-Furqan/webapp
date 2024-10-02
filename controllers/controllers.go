package controllers

import (
	"net/http"
	"strconv"

	"github.com/E-Furqan/webapp/data"
	"github.com/E-Furqan/webapp/database"
	"github.com/gin-gonic/gin"
)

func Getdata(c *gin.Context) {
	var students []data.Student_info
	if err := database.DB.Find(&students).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, students)
}

func Create_student(c *gin.Context) {
	var student data.Student_info
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&student)
	c.JSON(http.StatusCreated, student)
}

func Delete_student(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var student data.Student_info
	database.DB.First(&student, id)
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	database.DB.Delete(&student)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})

}

func Update_info(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var student data.Student_info
	database.DB.First(&student, id)
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var updatastudent data.Student_info
	err1 := c.ShouldBindJSON(&updatastudent)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}

	if err := database.DB.Model(&student).Updates(updatastudent).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update student"})
		return
	}
	c.JSON(http.StatusAccepted, student)
}
