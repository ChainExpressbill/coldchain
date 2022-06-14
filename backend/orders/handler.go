package orders

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// @Summary 		주문 조회
// @Description 주문 조회 ( page, size, startDate, endDate 필수** )
// @Tags 				Order
// @Accept 			application/json;charset=UTF-8
// @Produce 		application/json;charset=UTF-8
// @Param 			OrderSearchParams query OrderSearchParams true "주문 조회 조건"
// @Success 		200 {object} GetOrdersResponse
// @Failure 		400 {object} map[string]string
// @Failure 		500 {object} map[string]string
// @Router 			/orders [get]
func Orders(c *fiber.Ctx) error {
	params := new(OrderSearchParams)

	if err := c.QueryParser(params); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	orders := OrdersService(params)
	return c.Status(fiber.StatusOK).JSON(orders)
}

// @Summary 		주문 상세 조회
// @Description 주문 상세 조회
// @Tags 				Order
// @Accept 			application/json;charset=UTF-8
// @Produce 		application/json;charset=UTF-8
// @Param 			id 	path string true "order id"
// @Success 		200 {object} ent.Order
// @Failure 		400 {object} map[string]string
// @Failure 		500 {object} map[string]string
// @Router 			/orders/{id} [get]
func OrderDetail(c *fiber.Ctx) error {
	param := c.Params("id")

	if param == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// 입력 문자열이 정수 형식이 아닌 경우 함수는 0을 반환합니다.
	id, err := strconv.Atoi(param)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	order := OrderDetailService(id)
	return c.Status(fiber.StatusOK).JSON(order)
}

// @Summary 		주문 등록
// @Description 주문 등록
// @Tags 				Order
// @Accept 			application/json;charset=UTF-8
// @Produce 		application/json;charset=UTF-8
// @Param 			OrderRequestBody body OrderRequestBody true "주문 등록에 대한 form 데이터. id, oid 필요 없음"
// @Success 		200 {object} string
// @Failure 		400 {object} map[string]string
// @Failure 		500 {object} map[string]string
// @Router 			/orders [post]
func OrderCreate(c *fiber.Ctx) error {
	body := new(OrderRequestBody)

	if err := c.BodyParser(body); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	OrderCreateService(body)
	return c.SendStatus(fiber.StatusOK)
}

// @Summary 		주문 수정
// @Description 주문 수정
// @Tags 				Order
// @Accept 			application/json;charset=UTF-8
// @Produce 		application/json;charset=UTF-8
// @Param 			OrderRequestBody body OrderRequestBody true "주문 수정에 대한 form 데이터. accountId, oid 필요 없음"
// @Success 		200 {object} string
// @Failure 		400 {object} map[string]string
// @Failure 		500 {object} map[string]string
// @Router 			/orders [put]
func OrderUpdate(c *fiber.Ctx) error {
	body := new(OrderRequestBody)

	if err := c.BodyParser(body); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	OrderUpdateService(body)
	return c.SendStatus(fiber.StatusOK)
}
