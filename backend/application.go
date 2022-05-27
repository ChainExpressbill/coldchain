package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ChainExpressbill/coldchain/account"
	"github.com/ChainExpressbill/coldchain/dashboard"
	"github.com/ChainExpressbill/coldchain/database"
	"github.com/ChainExpressbill/coldchain/middlewares"
	"github.com/ChainExpressbill/coldchain/orders"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"github.com/qinains/fastergoding"
)

// env file loading
func loadEnvironments() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func useHotReloading() {
	STAGE := os.Getenv("STAGE")

	if STAGE != "prod" {
		fastergoding.Run() // hot reload
	}
}

func main() {
	loadEnvironments()
	useHotReloading()

	client := database.GetInstance()
	defer client.Close()

	database.MigrationDatabase()

	app := fiber.New()
	app.Use(recover.New())
	app.Use(middlewares.Cors())

	app.Post("/login", account.Login)
	app.Post("/logout", account.Logout)
	app.Post("/join", account.Join)

	dashboardRouter := app.Group("/dashboard")
	dashboardRouter.Get("/summary/last-month", dashboard.LastMonth)
	dashboardRouter.Get("/summary/today", dashboard.Today)
	dashboardRouter.Get("/summary/charts/:type", dashboard.Charts)

	ordersRouter := app.Group("/orders")
	ordersRouter.Get("/", orders.Orders)
	ordersRouter.Get("/:id", orders.OrderDetail)
	ordersRouter.Post("/", orders.OrderCreate)
	ordersRouter.Put("/", orders.OrderUpdate)

	port := fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))
	log.Fatalln(app.Listen(port))
}
