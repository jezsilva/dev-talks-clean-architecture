package core_use_case_repository_interface

import "devtalks/core/entity"

type OrderRepository interface {
	Create(order core_entity.Order) (core_entity.Order, error)
	GetByCodeExternal(codeExternal string) (core_entity.Order, error)
}
