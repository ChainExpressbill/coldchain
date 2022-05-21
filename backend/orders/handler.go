package orders

import "github.com/gofiber/fiber/v2"

func Orders(c *fiber.Ctx) error {
	return c.SendString("Orders")
}

func OrderDetail(c *fiber.Ctx) error {
	return c.SendString("OrderDetail")
}

func OrderCreate(c *fiber.Ctx) error {
	return c.SendString("OrderCreate")
}

func OrderUpdate(c *fiber.Ctx) error {
	return c.SendString("OrderUpdate")
}
