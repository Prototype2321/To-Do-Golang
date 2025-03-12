package main

import (
	"net/http"
	"strconv"

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
	router.GET("/tasks/:id", getTaskById)
	router.POST("/tasks", postTasks)

	router.Run("localhost:8080")
}

func getTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
}

func postTasks(c *gin.Context) {
	var newTask task

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	tasks = append(tasks, newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
}

func getTaskById(c *gin.Context) {
	id := c.Param("id")

	taskId, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "id must be integer"})
		return
	}

	for _, t := range tasks {
		if t.ID == taskId {
			c.IndentedJSON(http.StatusOK, t)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
}
