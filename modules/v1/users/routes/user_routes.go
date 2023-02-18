package routes

import (
	database "FastStarter/infrastructures/databases"
	usersControllerV1 "FastStarter/modules/v1/users/interfaces/controllers"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(router *fiber.App, db database.Database) *fiber.App {
	usersControllerV1 := usersControllerV1.UserController(db)

	api := router.Group("/api/v1")
	api.Get("/", usersControllerV1.Hello)

	return router
}
