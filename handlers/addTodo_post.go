package handlers

import (
	"hello_go/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddTodoEndPoint(data model.TodoList) gin.HandlerFunc {
	return func(c *gin.Context) {
		receivedData := model.Todo{}
		c.Bind(&receivedData)

		data.Add(receivedData)
		c.Status(http.StatusNoContent)
	}
}
