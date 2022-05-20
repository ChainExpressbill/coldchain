package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/qinains/fastergoding"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fastergoding.Run() // hot reload
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	port := fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))
	log.Fatalln(app.Listen(port))
}
