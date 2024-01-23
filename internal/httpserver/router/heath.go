package router

import "github.com/gin-gonic/gin"

func (r *RouterManager) RegisterHealthRoutes(path string) {
	r.handler.GET(path, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
