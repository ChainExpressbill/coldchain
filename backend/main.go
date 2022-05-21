package main

import (
	"fmt"
	"log"
	"os"

	"entgo.io/ent/examples/fs/ent"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/qinains/fastergoding"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_USER := os.Getenv("DB_USER")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	connectionUrl := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", DB_HOST, DB_PORT, DB_USER, DB_NAME, DB_PASSWORD)
	client, err := ent.Open("postgres", connectionUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	fastergoding.Run() // hot reload
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	port := fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))
	log.Fatalln(app.Listen(port))
}
