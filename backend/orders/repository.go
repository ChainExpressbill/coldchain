package orders

import (
	"context"

	"github.com/ChainExpressbill/coldchain/database"
	"github.com/ChainExpressbill/coldchain/ent"
	"github.com/ChainExpressbill/coldchain/ent/order"
)

var ctx = context.TODO()

func OrderCountByOrdererAndReceiver(params *OrderSearchParams) int {
	result := database.GetInstance().Order.
		Query().
		Where(
			order.OrdererContains(params.Orderer),
			order.ReceiverContains(params.Receiver),
			order.CreatedAtGTE(params.StartDate),
			order.CreatedAtLTE(params.EndDate),
		).
		CountX(ctx)

	return result
}

func FindAllByOrdererAndReceiver(params *OrderSearchParams) []*ent.Order {
	/**
	** @see https://entgo.io/docs/crud#field-selection
	**/
	result := database.GetInstance().Order.
		Query().
		Offset(params.Page-1).
		Limit(params.Size).
		Order(ent.Desc(order.FieldCreatedAt)).
		Where(
			order.OrdererContains(params.Orderer),
			order.ReceiverContains(params.Receiver),
			order.CreatedAtGTE(params.StartDate),
			order.CreatedAtLTE(params.EndDate),
		).
		AllX(ctx)

	return result
}

func FindById(id int) *ent.Order {
	result := database.GetInstance().Order.
		Query().
		Where(order.ID(id)).
		OnlyX(ctx)

	return result
}

func CreateOrder(body *OrderRequestBody) {
	database.GetInstance().Order.
		Create().
		SetOrderer(body.Orderer).
		SetReceiver(body.Receiver).
		SetDrugName(body.DrugName).
		SetDrugStandard(body.DrugStandard).
		SetQuantity(body.Quantity).
		SetRegisterName(body.RegisterName).
		SetStorageCondition(body.StorageCondition).
		SetManagerID(body.AccountId).
		SaveX(ctx)
}

func UpdateOrder(body *OrderRequestBody) {
	database.GetInstance().Order.
		UpdateOneID(body.Id).
		SetOrderer(body.Orderer).
		SetReceiver(body.Receiver).
		SetDrugName(body.DrugName).
		SetDrugStandard(body.DrugStandard).
		SetQuantity(body.Quantity).
		SetRegisterName(body.RegisterName).
		SetStorageCondition(body.StorageCondition).
		SaveX(ctx)
}
