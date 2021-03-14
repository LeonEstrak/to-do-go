package main

import (
	"hello_go/model"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	todoData := model.New()

	//NEED TO PUT CORS POLICY ***BEFORE*** registering any routes
	router.Use(cors.Default())

	todoData.Add(model.Todo{ID: "1", Completed: false, Message: "message1 GO"})
	todoData.Add(model.Todo{ID: "2", Completed: false, Message: "message2 GO"})
	todoData.Add(model.Todo{ID: "3", Completed: false, Message: "message3 GO"})

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "HomePage of Aniket",
		})
	})

	router.GET("/getTodo", func(c *gin.Context) {
		c.JSON(http.StatusOK, todoData.GetAll())
	})

	router.Run()
}
