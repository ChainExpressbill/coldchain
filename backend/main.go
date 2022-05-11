package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/qinains/fastergoding"
)

const port string = ":28342"

func main() {
	fastergoding.Run() // hot reload
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatalln(app.Listen(port))
}
