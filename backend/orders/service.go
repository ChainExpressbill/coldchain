package orders

import (
	"github.com/ChainExpressbill/coldchain/ent"
	"github.com/google/uuid"
)

func OrdersService(params *OrderSearchParams) []*ent.Order {
	return FindAll(params)
}

func OrderDetailService(id uuid.UUID) *ent.Order {
	return FindById(id)
}

func OrderCreateService(body *OrderRequestBody) {
	CreateAccount(body)
}

func OrderUpdateService(id uuid.UUID) {

}
