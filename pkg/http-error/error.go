package error

import "github.com/gofiber/fiber/v2"

func PageNotFound() func(c *fiber.Ctx) {
	return func(c *fiber.Ctx) {
		c.Status(404).Render("errors/error_404", fiber.Map{
			"Title": "Error 404",
		})
	}
}
