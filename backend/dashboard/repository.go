package dashboard

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/ChainExpressbill/coldchain/database"
	"github.com/ChainExpressbill/coldchain/ent/order"
)

var ctx = context.TODO()

const TIMESTAMP_COLUMN = "TO_TIMESTAMP(TO_CHAR(created, 'YYYY-MM-DD'), 'YYYY-MM-DD')"

func OrderCountByCreated(startDate, endDate time.Time) int {
	result := database.GetInstance().Order.
		Query().
		Where(
			order.CreatedAtGTE(startDate),
			order.CreatedAtLTE(endDate),
		).
		CountX(ctx)

	return result
}

func CountGroupByOrdererAndCreated(startDate, endDate time.Time) []string {
	result := database.GetInstance().Order.
		Query().
		Where(
			order.CreatedAtGTE(startDate),
			order.CreatedAtLTE(endDate),
		).
		GroupBy(order.FieldOrderer).
		StringsX(ctx)

	return result
}

func (cr *ChartsResponse) OrderCountGroupByCreated(startDate, endDate time.Time) {
	database.GetInstance().Order.
		Query().
		Where(
			order.CreatedAtGTE(startDate),
			order.CreatedAtLTE(endDate),
		).
		Modify(func(s *sql.Selector) {
			s.Select(
				sql.As(TIMESTAMP_COLUMN, "timestamp"),
				sql.As(sql.Count("*"), "count"),
			).
				GroupBy(TIMESTAMP_COLUMN).
				OrderBy(sql.Asc(TIMESTAMP_COLUMN))
		}).
		ScanX(ctx, cr)
}

func (cr *ChartsResponse) OrdererCountGroupByCreated(startDate, endDate time.Time) {
	database.GetInstance().Order.
		Query().
		Where(
			order.CreatedAtGTE(startDate),
			order.CreatedAtLTE(endDate),
		).
		Modify(func(s *sql.Selector) {
			s.Select(
				sql.As(TIMESTAMP_COLUMN, "timestamp"),
				sql.As("orderer", "orderer"),
				sql.As(sql.Count("orderer"), "count"),
			).
				GroupBy(TIMESTAMP_COLUMN, "orderer").
				OrderBy(sql.Asc(TIMESTAMP_COLUMN))
		}).
		ScanX(ctx, cr)
}
