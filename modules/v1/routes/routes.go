package routes

import "github.com/gofiber/fiber/v2"

func ParseTmpl(router *fiber.App) *fiber.App {
	router.Static("/", "./public")
	// router.Static("/static", "./public/static")
	// router.Static("/assets", "./public/assets")
	// router.Static("/favicon.ico", "./public/favicon.ico")
	// router.Static("/robots.txt", "./public/robots.txt")
	// router.Static("/sitemap.xml", "./public/sitemap.xml")
	return router
}

func NewRouter(router *fiber.App, db map[string]interface{}) *fiber.App {

	api := router.Group("/api/v1")
	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

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
