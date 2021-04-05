package handlers

import (
	"fmt"
	"hello_go/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddTodoEndPoint(data *model.TodoList) gin.HandlerFunc {
	return func(c *gin.Context) {
		receivedData := model.Todo{}
		c.Bind(&receivedData)
		fmt.Println(receivedData)

		data.Add(receivedData)
		fmt.Println(data)
		c.Status(http.StatusNoContent)
	}
}
