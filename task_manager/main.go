package main

import (
	"log"
	"task_manager/controllers"
	"task_manager/data"
	"task_manager/router"
)

func main() {
	// Initialize memory store service
	taskService := data.NewTaskService()

	// Initialize controller with service dependency
	taskController := controllers.NewTaskController(taskService)

	// Setup routes
	r := router.SetupRouter(taskController)

	// Run server on localhost:8080
	log.Println("Server starting on http://localhost:8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}