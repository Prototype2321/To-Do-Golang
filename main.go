package main



import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type task struct{
	ID 						int				`json:"id"`
	Title 				string		`json:"title"`
	Description 	string		`json:"description"`
	Status 				bool		 	`json:"status"`
}

var tasks = []task{
	{ID: 1, Title: "Belajar Go", Description: "Percobaan Pertama", Status: false},
	{ID: 2, Title: "Belajar Rest API", Description: "Percobaan Kedua", Status: false},
	{ID: 3, Title: "Belajar Gin", Description: "Percobaan Ketiga", Status: false},
}

func main() {
	router := gin.Default()
	router.GET("/tasks", getTask)
	router.GET("/tasks/:id", getTaskByID)
	router.POST("/tasks", postTasks)

	router.Run("localhost:8080")
}

func getTask(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
}

func postTasks(c *gin.Context){
	var newTask task

	if err:= c.BindJSON(&newTask) ; err !=nil{
		return
	}
	tasks = append(tasks, newTask)
	c.IndentedJSON(http.StatusCreated, newTask) 
}

func getTaskByID(c *gin.Context){
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "id must be integer"})
		return
	}
	for _, a := range tasks {
		
		if a.ID == idInt {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found" })
}

