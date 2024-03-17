package controllers

import (
	"go-rest-api/database"
	"go-rest-api/models"
	"go-rest-api/utils"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	db := database.GetDatabase()
	var order models.Order

	err := c.ShouldBindJSON(&order)
	if err != nil {
		utils.Respfailed("Json хөрвүүлэх үед алдаа гарлаа !!! ", c, err.Error())
		return
	}

	err = db.Create(&order).Error
	if err != nil {
		utils.Respfailed("Захиалга үүсгэж чадсангүй !!! ", c, err.Error())
		return
	}

	utils.RespSuccess(order, "", c)
}

func GetAllOrders(c *gin.Context) {
	db := database.GetDatabase()
	var p []models.Order
	err := db.Find(&p).Error

	if err != nil {
		utils.Respfailed("Cannot find User with id: !!! ", c, err.Error())
		return
	}
	utils.RespSuccess(p, "", c)
}
