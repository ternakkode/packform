package router

import (
	"github.com/gin-gonic/gin"
)

type RouterManager struct {
	handler *gin.Engine
}

func NewRouteManager(
	handler *gin.Engine,
) *RouterManager {
	return &RouterManager{
		handler: handler,
	}
}

func (r *RouterManager) Register() {
	r.RegisterHealthRoutes("health")
}
