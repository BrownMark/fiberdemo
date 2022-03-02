package main

import (
	"fiberexample/example/routers"
	"fiberexample/example/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/helmet/v2"
	"log"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
func main() {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(helmet.New())
	app.Use(logger.New())

	utils.ConfigureSwagger(app)
	routers.ConfigureTest(app)

	log.Fatal(app.Listen(":3000"))
}
