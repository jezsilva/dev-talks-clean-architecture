package controller

import (
	"devtalks/core/entity"
	"devtalks/core/use_case"
	"encoding/json"
)

type OrderController struct {
	CreateOrderUseCase entity_use_case.CreateOrderUseCase
}

func (controller OrderController) CreateOrder(body []byte) (interface{}, error) {
	var order core_entity.Order

	err := json.Unmarshal(body, &order)
	if err != nil {
		return nil, err
	}

	order, err = controller.CreateOrderUseCase.Execute(order)

	return order, err
}
