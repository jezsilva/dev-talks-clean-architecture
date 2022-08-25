package repository_memory

import "devtalks/core/entity"

type OrderRepository struct {
	Orders []core_entity.Order
}

func (repository *OrderRepository) Create(order core_entity.Order) (core_entity.Order, error) {
	repository.Orders = append(repository.Orders, order)

	return order, nil
}

func (repository *OrderRepository) GetByCodeExternal(codeExternal string) (core_entity.Order, error) {
	for _, order := range repository.Orders {
		if order.CodeExternal == codeExternal {
			return order, nil
		}
	}

	return core_entity.Order{}, nil
}
