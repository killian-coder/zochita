package handlers

import (
	"net/http"

	"example/todo-rest-api/models"

	"github.com/gin-gonic/gin"
)

var todos = []models.Todo{
	{ID: "1", Item: "clean room ", Completed: false},
	{ID: "2", Item: "Read Book ", Completed: false},
	{ID: "3", Item: "Record Video ", Completed: false},
}

func GetTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func AddTodo(context *gin.Context) {
	var newTodo models.Todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}
