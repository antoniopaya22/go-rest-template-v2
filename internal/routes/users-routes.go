package routes

import (
	"ai4am.com/go-restapi/internal/controllers"
	"github.com/gin-gonic/gin"
)

func SetUsersRoutes(app *gin.RouterGroup)  {
	app.GET("/users/:id", controllers.GetUserById)
	app.GET("/users", controllers.GetAllUsers)
	app.POST("/users", controllers.CreateUser)
	app.PUT("/users/:id", controllers.UpdateUser)
	app.DELETE("/users/:id", controllers.DeleteUser)
}

