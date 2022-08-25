package entity_use_case

import (
	"devtalks/core/entity"
	"devtalks/core/use_case/repository_interface"
)

type CreateOrderUseCase struct {
	OrderRepository core_use_case_repository_interface.OrderRepository
	OrderService    core_entity.OrderExistsService
}

func (os CreateOrderUseCase) Execute(order core_entity.Order) (core_entity.Order, error) {
	err := order.Validate()
	if err != nil {
		return core_entity.Order{}, err
	}

	orderExists, err := os.OrderRepository.GetByCodeExternal(order.CodeExternal)

	if err != nil {
		return core_entity.Order{}, err
	}

	err = os.OrderService.CreateOrderValidate(orderExists)
	if err != nil {
		return core_entity.Order{}, err
	}

	newOrder, err := os.OrderRepository.Create(order)

	if err != nil {
		return core_entity.Order{}, err
	}

	return newOrder, nil
}
