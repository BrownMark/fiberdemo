package middlewares

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func CustomLoggingMiddleware(c *fiber.Ctx) error {
	fmt.Println("called")
	return c.Next()
}
