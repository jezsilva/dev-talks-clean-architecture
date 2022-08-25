package main

import (
	"devtalks/controller"
	"devtalks/core/entity"
	"devtalks/core/use_case"
	repository_memory "devtalks/repository/memory"
	"devtalks/web"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	orderRepository := repository_memory.OrderRepository{}

	orderService := core_entity.OrderExistsService{}
	createOrderUseCase := entity_use_case.CreateOrderUseCase{
		OrderRepository: &orderRepository,
		OrderService:    orderService,
	}
	orderController := controller.OrderController{
		CreateOrderUseCase: createOrderUseCase,
	}
	routeGin := web.GinRoute{
		Engine:          r,
		OrderController: orderController,
	}
	routeGin.Inicialize()

	r.Run()
}
