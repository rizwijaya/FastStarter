package handler

import "github.com/gofiber/fiber/v2"

func (u *usersHandler) Hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}
