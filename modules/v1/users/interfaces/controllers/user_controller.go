package controllers

import "github.com/gofiber/fiber/v2"

func (uc *UsersController) Hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}
