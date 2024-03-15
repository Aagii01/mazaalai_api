package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rromulos/go-rest-api/database"
	"github.com/rromulos/go-rest-api/models"
)

func Login(c *gin.Context) {
	db := database.GetDatabase()

	var p models.User

	err := c.ShouldBindJSON(&p)

	fmt.Printf("gmail: %v \n", p.Email)
	fmt.Printf("password: %v \n", p.Password)

	if err != nil {
		c.JSON(400, gin.H{
			"Error": "Cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Find(&p).Error
	if err != nil {
		c.JSON(400, gin.H{
			"Error": "Error Finding the User: " + err.Error(),
		})
		return
	}

	var name string
	name = "amjilttati"

	c.JSON(200, name)
}
