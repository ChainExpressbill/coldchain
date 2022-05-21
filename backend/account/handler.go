package account

import "github.com/gofiber/fiber/v2"

func Login(c *fiber.Ctx) error {
	return c.SendString("Login")
}

func Logout(c *fiber.Ctx) error {
	return c.SendString("Logout")
}

func Join(c *fiber.Ctx) error {
	return c.SendString("Join")
}
