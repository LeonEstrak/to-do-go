package main

import (
	"hello_go/handlers"
	"hello_go/model"

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

	// Use '&' to pass the pointer reference of the variable `todoData`
	router.GET("/todo", handlers.GetTodoEndPoint(&todoData))
	router.POST("/todo", handlers.AddTodoEndPoint(&todoData))
	router.PUT("/todo", handlers.UpdateTodoEndPoint(&todoData))
	router.DELETE("/todo", handlers.DeleteTodoEndPoint(&todoData))
	router.Run()
}
