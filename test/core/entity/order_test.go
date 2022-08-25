package test_entity

import (
	core_entity "devtalks/core/entity"
	"testing"
)

func TestValidateInvalidExternalCode(t *testing.T) {
	order := core_entity.Order{}

	err := order.Validate()
	if err == nil {
		t.Errorf("Deve retornar erro")
	}
}

func TestValidateInvalidTotalAmount(t *testing.T) {
	order := core_entity.Order{
		CodeExternal: "1",
	}

	err := order.Validate()
	if err == nil {
		t.Errorf("Deve retornar erro")
	}
}

func TestValidate(t *testing.T) {
	order := core_entity.Order{
		CodeExternal: "1",
		TotalAmount:  1,
	}

	err := order.Validate()
	if err != nil {
		t.Errorf("Erro inesperado")
	}
}
