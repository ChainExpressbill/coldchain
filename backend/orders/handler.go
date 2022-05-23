package orders

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
	oid := c.Params("id")

	if oid == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	fmt.Printf("Oid: %s", oid)

	id, err := uuid.ParseBytes([]byte(oid))

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	order := OrderDetailService(id)
	fmt.Printf("%#v\n", order)
	return c.SendString("OrderDetail")
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

	return c.SendString("OrderUpdate")
}
