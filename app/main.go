package main

import (
	"FastStarter/infrastructures/config"
	database "FastStarter/infrastructures/databases"
	routesAPIV1 "FastStarter/modules/v1/users/routes"
	routesTMPLV1 "FastStarter/public/v1/handler/v1/routes"
	"errors"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"
)

func NewRouting() (*fiber.App, database.Database) {
	database := database.NewDatabase()
	engine := html.New("./public/v1/templates", ".html")
	router := fiber.New(fiber.Config{
		Views: engine,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			err = ctx.Status(code).Render(fmt.Sprintf("errors/%d", code), fiber.Map{
				"Title": "Error " + fmt.Sprintf("%d", code),
			}, "layouts/main")
			//Render(fmt.Sprintf("./public/v1/templates/errors/%d.html", code))
			if err != nil {
				ctx.Status(fiber.StatusInternalServerError).Render("errors/500.html", fiber.Map{
					"Title": "Error " + fmt.Sprintf("%d", code),
				}, "layouts/main")
			}
			return nil
		},
	})

	router.Use(cors.New())

	return router, database
}

func main() {
	config, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	router, db := NewRouting()
	router = routesAPIV1.NewRouter(router, db)
	router = routesTMPLV1.NewRouter(router, db)

	router.Listen(config.App.Url + ":" + config.App.Port)
}
