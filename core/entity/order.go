package core_entity

import "errors"

type OrderItems struct {
	CodeExternal string
	UnitPrice    float64
	TotalPrice   float64
	Quantity     float64
	Description  string
}

type Order struct {
	Id           string
	CodeExternal string
	TotalAmount  float64
	OrderItems   []OrderItems
}

func (o Order) Validate() error {
	if len(o.CodeExternal) <= 0 {
		return errors.New("Código externo deve ser informado")
	}

	if o.TotalAmount <= 0 {
		return errors.New("Valor total deve ser superior a zero")
	}

	return nil
}
