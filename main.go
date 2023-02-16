package main

import (
	"FastStarter/app/config"
	database "FastStarter/app/databases"
	routesV1 "FastStarter/modules/v1/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"
)

func new() (*fiber.App, map[string]interface{}) {
	database, err := database.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}

	engine := html.New("./public/templates", ".html")
	router := fiber.New(fiber.Config{
		Views: engine,
	})

	router.Use(cors.New())

	return router, database
}

func main() {
	config, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	router := routesV1.NewRouter(new())
	router.Listen(config.App.Url + ":" + config.App.Port)
}
