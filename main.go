package main

import (
	"example/todo-rest-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/todos", handlers.GetTodos)
	// router.GET("/todos/:id", handlers.GetTodo)
	// router.PATCH("/todos/:id", handlers.ToggleTodoStatus)
	router.GET("/todo/get", handlers.GetDataFromDB)
	router.POST("/todos", handlers.AddTodo)
	router.Run("localhost:9090")
}
