package account

import (
	"github.com/gofiber/fiber/v2"
)

// @Description Chain expressbill 관리자 로그인
// @Summary 	Chain expressbill 관리자 로그인
// @Tags 			Account
// @Accept 		application/json;charset=UTF-8
// @Produce 	application/json;charset=UTF-8
// @Param 		loginParams body LoginParams true "id, password"
// @Success 	200 {object} ent.Account
// @Failure 	400 {object} map[string]string
// @Failure 	500 {object} map[string]string
// @Router 		/login [post]
func Login(c *fiber.Ctx) error {
	params := new(LoginParams)

	if err := c.BodyParser(params); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	user, err2 := LoginService(params.Id, params.Password, c)
	if err2 != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"code": "500", "message": "login failed"})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

// @Description Chain expressbill 관리자 로그아웃
// @Summary 	Chain expressbill 관리자 로그아웃
// @Tags 			Account
// @Accept 		application/json;charset=UTF-8
// @Produce 	application/json;charset=UTF-8
// @Param 		logoutParam body LoginParams true "id"
// @Success 	200 {object} string
// @Failure 	400 {object} map[string]string
// @Failure 	500 {object} map[string]string
// @Router 		/logout [post]
func Logout(c *fiber.Ctx) error {
	params := new(LoginParams)

	if err := c.BodyParser(params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": "400", "message": "id invalid"})
	}

	LogoutService(params.Id)
	return c.SendStatus(fiber.StatusOK)
}

// @Description Chain expressbill 관리자 회원가입
// @Summary 	Chain expressbill 관리자 회원가입
// @Tags 			Account
// @Accept 		application/json;charset=UTF-8
// @Produce 	application/json;charset=UTF-8
// @Param 		JoinParams body JoinParams true "id, password, name, emailAddress"
// @Success 	200 {object} string
// @Failure 	400 {object} map[string]string
// @Failure 	500 {object} map[string]string
// @Router 		/join [post]
func Join(c *fiber.Ctx) error {
	params := new(JoinParams)

	if err := c.BodyParser(params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": "400", "message": "join params invalid"})
	}

	JoinService(params)
	return c.SendStatus(fiber.StatusOK)
}
