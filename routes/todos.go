package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addTodoRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/todos")

	users.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "todos")
	})
}
