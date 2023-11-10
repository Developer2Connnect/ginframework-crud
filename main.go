package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Initialize Model
	model := NewModel()

	// Initialize Controller
	controller := NewController(model)

	// Routes
	router.GET("/todos", controller.GetTodos)
	router.GET("/todos/:id", controller.GetTodo)
	router.POST("/todos", controller.AddTodo)
	router.PUT("/todos/:id", controller.UpdateTodo)
	router.DELETE("/todos/:id", controller.DeleteTodo)

	// Run the server
	router.Run(":8080")
}
