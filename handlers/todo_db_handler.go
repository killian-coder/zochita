package handlers

import (
	"net/http"

	"example/todo-rest-api/config"
	"example/todo-rest-api/models"

	"github.com/gin-gonic/gin"
)

func GetDataFromDB(context *gin.Context) {

	db, err := config.DatabaseConnect()
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to connect to the database"})
		return
	}
	defer db.Close()

	var todosFromDB []models.Todo
	rows, err := db.Query("SELECT id, item, completed FROM todos")
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to execute query"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var t models.Todo
		if err := rows.Scan(&t.ID, &t.Item, &t.Completed); err != nil {
			context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to find todo"})
			return
		}
		todosFromDB = append(todosFromDB, t)
	}

	if err = rows.Err(); err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to iterate over rows"})
		return
	}

	context.IndentedJSON(http.StatusOK, todosFromDB)
}
