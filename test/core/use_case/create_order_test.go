package test_use_case

import (
	core_entity "devtalks/core/entity"
	"devtalks/core/use_case"
	"errors"
	"testing"
)

func TestCreateOrderInvalidExternalCode(t *testing.T) {
	createOrder := entity_use_case.CreateOrderUseCase{}
	order := core_entity.Order{}

	_, err := createOrder.Execute(order)
	if err == nil {
		t.Errorf("é esperado um erro")
	}
}

func TestCreateOrderGetOrderError(t *testing.T) {
	orderRepository := mockOrderRepository{
		success: core_entity.Order{},
		error:   errors.New("Error"),
	}
	createOrder := entity_use_case.CreateOrderUseCase{
		OrderRepository: orderRepository,
	}
	order := core_entity.Order{
		CodeExternal: "1",
		TotalAmount:  1,
	}

	_, err := createOrder.Execute(order)
	if err == nil {
		t.Errorf("é esperado um erro")
	}
}

func TestCreateOrderGetOrderExists(t *testing.T) {
	orderRepository := mockOrderRepository{
		success: core_entity.Order{
			Id: "1",
		},
		error: nil,
	}
	orderService := core_entity.OrderExistsService{}
	createOrder := entity_use_case.CreateOrderUseCase{
		OrderRepository: orderRepository,
		OrderService:    orderService,
	}
	order := core_entity.Order{
		CodeExternal: "1",
		TotalAmount:  1,
	}

	_, err := createOrder.Execute(order)
	if err == nil {
		t.Errorf("é esperado um erro")
	}
}

func TestCreateOrderCreateOrderError(t *testing.T) {
	orderRepository := mockOrderRepository{
		success: core_entity.Order{
			Id: "1",
		},
		error: errors.New("Error"),
	}
	orderService := core_entity.OrderExistsService{}
	createOrder := entity_use_case.CreateOrderUseCase{
		OrderRepository: orderRepository,
		OrderService:    orderService,
	}
	order := core_entity.Order{
		CodeExternal: "1",
		TotalAmount:  1,
	}

	_, err := createOrder.Execute(order)
	if err == nil {
		t.Errorf("é esperado um erro")
	}
}

type mockOrderRepository struct {
	success core_entity.Order
	error   error
}

func (mock mockOrderRepository) Create(order core_entity.Order) (core_entity.Order, error) {
	return mock.success, mock.error
}

func (mock mockOrderRepository) GetByCodeExternal(codeExternal string) (core_entity.Order, error) {
	return mock.success, mock.error
}
