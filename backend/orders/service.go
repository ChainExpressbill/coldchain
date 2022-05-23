package orders

import (
	"math"

	"github.com/ChainExpressbill/coldchain/ent"
)

func OrdersService(params *OrderSearchParams) GetOrdersResponse {
	totalCount := OrderCountByOrdererAndReceiver(params)
	orderList := FindAllByOrdererAndReceiver(params)
	orders := GetOrdersResponse{
		OrderList:  orderList,
		Page:       params.Page,
		Size:       params.Size,
		TotalPage:  int(math.Ceil(float64(totalCount) / float64(params.Size))),
		TotalCount: totalCount,
	}

	return orders
}

func OrderDetailService(id int) *ent.Order {
	return FindById(id)
}

func OrderCreateService(body *OrderRequestBody) {
	CreateOrder(body)
}

func OrderUpdateService(body *OrderRequestBody) {
	UpdateOrder(body)
}
