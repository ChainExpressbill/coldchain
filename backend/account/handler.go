package account

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	params := new(LoginParams)

	if err := c.BodyParser(params); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	user := LoginService(params.Id, params.Password)
	fmt.Println(user)
	return c.Status(fiber.StatusOK).JSON(user)
}

func Logout(c *fiber.Ctx) error {
	return c.SendString("Logout")
}

func Join(c *fiber.Ctx) error {
	params := new(JoinParams)

	if err := c.BodyParser(params); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	fmt.Printf("Id: %s, Password: %s, Name: %s, EmailAddress: %s\n", params.Id, params.Password, params.Name, params.EmailAddress)
	JoinService(params)
	return c.SendStatus(fiber.StatusOK)
}
