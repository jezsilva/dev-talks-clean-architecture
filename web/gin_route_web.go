package web

import (
	"devtalks/adapter/web"
	"devtalks/controller"
	"github.com/gin-gonic/gin"
)

type GinRoute struct {
	Engine          *gin.Engine
	OrderController controller.OrderController
}

func (gin *GinRoute) Inicialize() {
	gin.Engine.POST("/order", adapter_web.GinHttpAdapter(gin.OrderController.CreateOrder))
}
