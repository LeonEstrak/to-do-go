package handlers

import (
	"hello_go/model"

	"github.com/gin-gonic/gin"
)

func GetTodoEndPoint(data model.TodoList) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, data.GetAll())
	}
}
