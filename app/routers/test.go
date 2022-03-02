package routers

import (
	"fiberexample/example/middlewares"
	"fiberexample/example/models"
	"fiberexample/example/utils"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func ConfigureTest(app *fiber.App) {
	group := app.Group("/test", middlewares.CustomLoggingMiddleware)

	group.Get("/", helloHandler)
	group.Post("/", postHandler)
	group.Put("/:testId", putHandler)
}

// @Summary      Says Hello123
// @Description  Not much123!
// @Tags         Test
// @Produce      json
// @Success      200  {object}  models.Test
// @Success      204
// @Router       /test/ [get]
func helloHandler(c *fiber.Ctx) error {
	p := models.Test{Name: "hello", Age: 1}

	return c.Status(fiber.StatusOK).JSON(p)
}

// @Summary      Sample Put
// @Description  Not much123!
// @Tags         Test
// @Accept       json
// @Produce      json
// @Param        testId   path      int  true  "Test Id"
// @Param        data   body models.Test true "the update data example"
// @Success      200  {object}  models.Test
// @Failure      400  {object}  utils.ErrorResponse
// @Router       /test/{testId} [put]
func putHandler(c *fiber.Ctx) error {
	testId, err := strconv.Atoi(c.Params("testId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("TestId supplied is not a number")
	}

	test := new(models.Test)
	test.TestId = &testId

	if err := c.BodyParser(&test); err != nil {
		return err
	}

	if errors := utils.ValidateStruct(*test); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	return c.Status(fiber.StatusOK).JSON(test)
}

// @Summary      Sample Post123
// @Description  Not much13!
// @Tags         Test
// @Accept       json
// @Produce      json
// @Param        data   body models.Test true "the create data example"
// @Success      200  {object}  models.Test
// @Failure      400  {object}  utils.ErrorResponse
// @Router       /test/ [post]
func postHandler(c *fiber.Ctx) error {
	test := new(models.Test)
	if err := c.BodyParser(&test); err != nil {
		return err
	}

	if errors := utils.ValidateStruct(*test); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	return c.Status(fiber.StatusOK).JSON(test)
}
