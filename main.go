package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "clean room ", Completed: false},
	{ID: "2", Item: "Read Book ", Completed: false},
	{ID: "3", Item: "Record Video ", Completed: false},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
	var newTodo todo
	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}
func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}
func toggleTodoStatus(context *gin.Context) {

	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)
}

func getTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

func getDataFromDB(context *gin.Context) {

	// Call the databaseConnect function to get the database connection
	db, err := databaseConnect()
	if err != nil {
		// Handle the error if the database connection fails
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to connect to the database"})
		return
	}
	defer db.Close()

	// Perform your database query here using the db object
	// For example, let's say you want to retrieve all todos
	var todosFromDB []todo
	rows, err := db.Query("SELECT id, item, completed FROM todos")
	if err != nil {
		// Handle the error if the query execution fails
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to execute query"})
		return
	}
	defer rows.Close()

	// Iterate over the rows and scan the result into the todosFromDB slice
	for rows.Next() {
		var t todo
		if err := rows.Scan(&t.ID, &t.Item, &t.Completed); err != nil {
			// Handle the error if the row scan fails
			context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to scan todo"})
			return
		}
		todosFromDB = append(todosFromDB, t)
	}

	// Check for errors from iterating over rows
	if err = rows.Err(); err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to iterate over rows"})
		return
	}

	// Send the retrieved todos back to the client
	context.IndentedJSON(http.StatusOK, todosFromDB)
}

func databaseConnect() (*sql.DB, error) {
	// Replace 'root:password@tcp(localhost:3306)/todo_db' with your actual username, password, host, and database name
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/todo_db")
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	// Check if the database is connected
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to the database: %v", err)
	}

	fmt.Println("Connected to the database successfully!")
	return db, nil
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.GET("/todo/get", getDataFromDB)
	router.POST("/todos", addTodo)
	router.Run("localhost:9090")
}
