package core_entity

import "errors"

type OrderExistsService struct {
}

func (orderExistsService OrderExistsService) CreateOrderValidate(existsOrder Order) error {
	if len(existsOrder.Id) > 0 {
		return errors.New("Order já existe!")
	}
	return nil
}
