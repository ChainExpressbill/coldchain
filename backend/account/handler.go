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

	user, err2 := LoginService(params.Id, params.Password, c)
	if err2 != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{"code": "500", "message": "login failed"})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func Logout(c *fiber.Ctx) error {
	params := new(LoginParams)

	if err := c.BodyParser(params); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	LogoutService(params.Id)
	return c.SendStatus(fiber.StatusOK)
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
