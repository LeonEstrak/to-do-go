package handlers

import (
	"fmt"
	"hello_go/model"

	"github.com/gin-gonic/gin"
)

func DeleteTodoEndPoint(data *model.TodoList) gin.HandlerFunc {
	return func(c *gin.Context) {
		receivedData := model.Todo{}
		c.Bind(&receivedData)

		data.Delete(receivedData)
		fmt.Println(data)
		c.Status(200)
	}
}
