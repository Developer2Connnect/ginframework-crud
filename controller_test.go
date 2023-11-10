// controller_test.go

package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCRUDOperations(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	// Initialize Model and Controller
	model := NewModel()
	controller := NewController(model)

	// Routes
	router.GET("/todos", controller.GetTodos)
	router.GET("/todos/:id", controller.GetTodo)
	router.POST("/todos", controller.AddTodo)
	router.PUT("/todos/:id", controller.UpdateTodo)
	router.DELETE("/todos/:id", controller.DeleteTodo)

	// Test AddTodo
	t.Run("AddTodo", func(t *testing.T) {
		// Prepare a new Todo
		newTodo := Todo{Title: "Test Todo", Completed: false}

		// Convert Todo to JSON
		jsonTodo, err := json.Marshal(newTodo)
		assert.NoError(t, err)

		// Perform HTTP POST request to add a new Todo
		req, err := http.NewRequest("POST", "/todos", bytes.NewBuffer(jsonTodo))
		assert.NoError(t, err)

		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		// Verify the HTTP response
		assert.Equal(t, http.StatusCreated, resp.Code)

		// Parse the response body
		var addedTodo Todo
		err = json.Unmarshal(resp.Body.Bytes(), &addedTodo)
		assert.NoError(t, err)

		// Verify that the Todo was added
		assert.Equal(t, newTodo.Title, addedTodo.Title)
		assert.False(t, addedTodo.Completed)
		assert.NotZero(t, addedTodo.ID)
	})

	// Test GetTodos
	t.Run("GetTodos", func(t *testing.T) {
		// Perform HTTP GET request to retrieve all Todos
		req, err := http.NewRequest("GET", "/todos", nil)
		assert.NoError(t, err)

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		// Verify the HTTP response
		assert.Equal(t, http.StatusOK, resp.Code)

		// Parse the response body
		var todos []Todo
		err = json.Unmarshal(resp.Body.Bytes(), &todos)
		assert.NoError(t, err)

		// Verify that the Todos slice is not empty
		assert.NotEmpty(t, todos)
	})

	// Test GetTodo
	t.Run("GetTodo", func(t *testing.T) {
		// Perform HTTP GET request to retrieve a specific Todo
		req, err := http.NewRequest("GET", "/todos/1", nil)
		assert.NoError(t, err)

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		// Verify the HTTP response
		assert.Equal(t, http.StatusOK, resp.Code)

		// Parse the response body
		var todo Todo
		err = json.Unmarshal(resp.Body.Bytes(), &todo)
		assert.NoError(t, err)

		// Verify that the correct Todo was retrieved
		assert.Equal(t, 1, todo.ID)
	})

	// Test UpdateTodo
	t.Run("UpdateTodo", func(t *testing.T) {
		// Prepare an updated Todo
		updatedTodo := Todo{Title: "Updated Todo", Completed: true}

		// Convert updated Todo to JSON
		jsonUpdatedTodo, err := json.Marshal(updatedTodo)
		assert.NoError(t, err)

		// Perform HTTP PUT request to update a Todo
		req, err := http.NewRequest("PUT", "/todos/1", bytes.NewBuffer(jsonUpdatedTodo))
		assert.NoError(t, err)

		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		// Verify the HTTP response
		assert.Equal(t, http.StatusOK, resp.Code)

		// Parse the response body
		var returnedTodo Todo
		err = json.Unmarshal(resp.Body.Bytes(), &returnedTodo)
		assert.NoError(t, err)

		// Verify that the Todo was updated
		assert.Equal(t, updatedTodo.Title, returnedTodo.Title)
		assert.True(t, returnedTodo.Completed)
	})

	// Test DeleteTodo
	t.Run("DeleteTodo", func(t *testing.T) {
		// Perform HTTP DELETE request to delete a Todo
		req, err := http.NewRequest("DELETE", "/todos/1", nil)
		assert.NoError(t, err)

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		// Verify the HTTP response
		assert.Equal(t, http.StatusOK, resp.Code)

		// Verify that the Todo was deleted
		assert.Empty(t, model.todos)
	})
}
