package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Cors(app *fiber.App) {
	// Or extend your config for customization
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // domain 정해지면 추가
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
}
