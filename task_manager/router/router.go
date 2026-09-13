package router

import (
	"task_manager/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(tc *controllers.TaskController) *gin.Engine {
	r := gin.Default()

	taskRoutes := r.Group("/tasks")
	{
		taskRoutes.GET("", tc.GetTasks)
		taskRoutes.GET("/:id", tc.GetTaskByID)
		taskRoutes.POST("", tc.CreateTask)
		taskRoutes.PUT("/:id", tc.UpdateTask)
		taskRoutes.DELETE("/:id", tc.DeleteTask)
	}

	return r
}