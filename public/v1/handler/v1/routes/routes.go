package routes

import (
	database "FastStarter/infrastructures/databases"

	"github.com/gofiber/fiber/v2"
)

func ParseTmpl(router *fiber.App) *fiber.App {
	router.Static("/", "./public")
	return router
}

func NewRouter(router *fiber.App, db database.Database) *fiber.App {
	views := router.Group("")
	views.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hello, World 👋!",
		})
	})

	views.Get("/layout", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hello, World 👋!",
		}, "layouts/main")
	})

	return router
}
