package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type task struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
	Completed   bool   `json:"completed"`
}

var tasks = []task{
	{ID: 1, Name: "Belajar Fundamental Go", Description: "Belajar dari basic bahasa golang", DueDate: "2025-01-01", Completed: true},
	{ID: 2, Name: "Belajar Fundamental Go Intermediate", Description: "Belajar golang rest api", DueDate: "2021-01-02", Completed: false},
	{ID: 3, Name: "Belajar Fundamental Go Advance", Description: "Belajar golang deployment", DueDate: "2021-01-05", Completed: false},
}

func main() {
	router := gin.Default()
	router.GET("/tasks", getTasks)

	router.Run("localhost:8080")
}

func getTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
}
