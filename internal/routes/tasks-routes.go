package routes

import (
	"ai4am.com/go-restapi/internal/controllers"
	"ai4am.com/go-restapi/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

func SetTasksRoutes(app *gin.RouterGroup)  {
	app.GET("/tasks/:id", controllers.GetTaskById)
	app.GET("/tasks", controllers.GetTasks)
	app.POST("/tasks", controllers.CreateTask)
	app.PUT("/tasks/:id", controllers.UpdateTask)
	app.DELETE("/tasks/:id", middlewares.AuthRequired(), controllers.DeleteTask)
}
