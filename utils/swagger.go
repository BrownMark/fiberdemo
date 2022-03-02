package utils

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func ConfigureSwagger(app *fiber.App) {
	swaggerGroup := app.Group("/swagger")
	swaggerGroup.Static("/doc.json", "./docs/swagger.json")
	swaggerGroup.Get("/*", swagger.HandlerDefault) // default
}
