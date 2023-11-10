package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	model *Model
}

func NewController(model *Model) *Controller {
	return &Controller{
		model: model,
	}
}

func (c *Controller) GetTodos(ctx *gin.Context) {
	todos := c.model.GetTodos()
	ctx.JSON(http.StatusOK, todos)
}

func (c *Controller) GetTodo(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	todo, found := c.model.GetTodoByID(id)
	if !found {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	ctx.JSON(http.StatusOK, todo)
}

func (c *Controller) AddTodo(ctx *gin.Context) {
	var newTodo Todo
	if err := ctx.BindJSON(&newTodo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	c.model.AddTodo(newTodo)
	ctx.JSON(http.StatusCreated, newTodo)
}

func (c *Controller) UpdateTodo(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedTodo Todo
	if err := ctx.BindJSON(&updatedTodo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if !c.model.UpdateTodoByID(id, updatedTodo) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	ctx.JSON(http.StatusOK, updatedTodo)
}

func (c *Controller) DeleteTodo(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if !c.model.DeleteTodoByID(id) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}
