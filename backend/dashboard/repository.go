package dashboard

import (
	"context"
	"time"

	"github.com/ChainExpressbill/coldchain/database"
	"github.com/ChainExpressbill/coldchain/ent/order"
)

var ctx = context.TODO()

func OrderCountByCreated(startDate, endDate time.Time) int {
	result := database.GetInstance().Order.
		Query().
		Where(
			order.CreatedGTE(startDate),
			order.CreatedLTE(endDate),
		).
		CountX(ctx)

	return result
}

func CountGroupByOrdererAndCreated(startDate, endDate time.Time) []string {
	result := database.GetInstance().Order.
		Query().
		Where(
			order.CreatedGTE(startDate),
			order.CreatedLTE(endDate),
		).
		GroupBy(order.FieldOrderer).
		StringsX(ctx)

	return result
}
