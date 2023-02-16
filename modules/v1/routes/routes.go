package routes

import "github.com/gofiber/fiber/v2"

func NewRouter(router *fiber.App, db map[string]interface{}) *fiber.App {

	api := router.Group("/api/v1")
	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	return router
}
