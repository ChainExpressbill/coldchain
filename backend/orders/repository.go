package orders

import (
	"context"

	"github.com/ChainExpressbill/coldchain/database"
	"github.com/ChainExpressbill/coldchain/ent"
	"github.com/ChainExpressbill/coldchain/ent/order"
	"github.com/google/uuid"
)

var ctx = context.TODO()

func FindAll(params *OrderSearchParams) []*ent.Order {
	result := database.GetInstance().Order.
		Query().
		WithManager().
		AllX(ctx)

	return result
}

func FindById(id uuid.UUID) *ent.Order {
	result := database.GetInstance().Order.
		Query().
		Where(order.Oid(id)).
		OnlyX(ctx)

	return result
}

func CreateAccount(body *OrderRequestBody) {
	database.GetInstance().Order.
		Create().
		SetOrderer(body.Orderer).
		SetReceiver(body.Receiver).
		SetDrugName(body.DrugName).
		SetDrugStandard(body.DrugStandard).
		SetQuantity(body.Quantity).
		SetRegisterName(body.RegisterName).
		SetStorageCondition(body.StorageCondition).
		SaveX(ctx)
}
