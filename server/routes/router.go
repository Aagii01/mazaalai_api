package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rromulos/go-rest-api/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		users := main.Group("users")
		{
			users.GET("/:id", controllers.GetUser)
			users.GET("/", controllers.GetAllUsers)
			users.POST("/", controllers.CreateUser)
			users.DELETE("/:id", controllers.DeleteUser)
			users.PUT("/", controllers.UpdateUser)
		}

		login := main.Group("login")
		{
			login.POST("/", controllers.Login)
		}
	}

	return router
}
