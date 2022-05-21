package dashboard

import "github.com/gofiber/fiber/v2"

func LastMonth(c *fiber.Ctx) error {
	return c.SendString("LastMonth")
}

func Today(c *fiber.Ctx) error {
	return c.SendString("Today")
}

func Charts(c *fiber.Ctx) error {
	return c.SendString("Charts")
}
