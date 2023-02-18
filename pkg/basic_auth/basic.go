package basic

import (
	"FastStarter/infrastructures/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func Auth(conf config.LoadConfig) func(*fiber.Ctx) error {
	return basicauth.New(basicauth.Config{
		Users: map[string]string{
			conf.BasicAuth.Username: conf.BasicAuth.Password,
		},
		Realm: "Restricted",
		Authorizer: func(user, pass string) bool {
			if user == conf.BasicAuth.Username && pass == conf.BasicAuth.Password {
				return true
			}
			return false
		},
		Unauthorized: func(c *fiber.Ctx) error {
			return c.SendFile("./public/templates/error/forbidden.html")
		},
		ContextUsername: "_user",
		ContextPassword: "_pass",
	},
	)
}
