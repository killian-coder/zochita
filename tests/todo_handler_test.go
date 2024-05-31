// tests/todo_handler_test.go
package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"example/todo-rest-api/handlers"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/todos", handlers.GetTodos)
	router.POST("/todos", handlers.AddTodo)
	return router
}

func TestGetTodos(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/todos", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var todos []handlers.Todo
	err := json.Unmarshal(w.Body.Bytes(), &todos)
	assert.Nil(t, err)
	assert.Equal(t, 3, len(todos))
}

func TestAddTodo(t *testing.T) {
	router := setupRouter()

	newTodo := handlers.Todo{ID: "4", Item: "test item", Completed: false}
	jsonValue, _ := json.Marshal(newTodo)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/todos", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var todo handlers.Todo
	err := json.Unmarshal(w.Body.Bytes(), &todo)
	assert.Nil(t, err)
	assert.Equal(t, "test item", todo.Item)
	assert.Equal(t, false, todo.Completed)
}
