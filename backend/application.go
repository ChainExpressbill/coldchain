package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ChainExpressbill/coldchain/account"
	"github.com/ChainExpressbill/coldchain/dashboard"
	"github.com/ChainExpressbill/coldchain/database"
	_ "github.com/ChainExpressbill/coldchain/docs"
	"github.com/ChainExpressbill/coldchain/middlewares"
	"github.com/ChainExpressbill/coldchain/orders"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
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

// @title Coldchain API
// @version 1.0
// @description This is coldchain app by chain-expressbill
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	loadEnvironments()
	useHotReloading()

	client := database.GetInstance()
	defer client.Close()

	database.MigrationDatabase()

	app := fiber.New()

	file, loggerConfig := middlewares.LoggerConfig()
	defer file.Close()

	app.Use(recover.New(), middlewares.CorsMiddleware(), logger.New(loggerConfig))

	app.Get("/swagger/*", swagger.HandlerDefault) // default

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
