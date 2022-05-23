package orders

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Orders(c *fiber.Ctx) error {
	params := new(OrderSearchParams)

	if err := c.QueryParser(params); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	fmt.Printf("%#v\n", params)
	orders := OrdersService(params)
	return c.Status(fiber.StatusOK).JSON(orders)
}

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
	fmt.Printf("%#v\n", order)
	return c.Status(fiber.StatusOK).JSON(order)
}

func OrderCreate(c *fiber.Ctx) error {
	body := new(OrderRequestBody)

	if err := c.BodyParser(body); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	fmt.Printf("%#v\n", body)
	OrderCreateService(body)
	return c.SendStatus(fiber.StatusOK)
}

func OrderUpdate(c *fiber.Ctx) error {
	body := new(OrderRequestBody)

	if err := c.BodyParser(body); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	fmt.Printf("%#v\n", body)
	OrderUpdateService(body)
	return c.SendStatus(fiber.StatusOK)
}
