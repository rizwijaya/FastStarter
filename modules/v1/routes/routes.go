package routes

import (
	database "FastStarter/app/databases"
	usersHandlerV1 "FastStarter/modules/v1/utilities/users/handler"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(router *fiber.App, db database.Database) *fiber.App {
	usersHandlerV1 := usersHandlerV1.Handler(db)

	api := router.Group("/api/v1")
	api.Get("/", usersHandlerV1.Hello)

	return router
}
