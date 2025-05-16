package handlers

import "github.com/gofiber/fiber/v2"

type WelcomeHandler struct{}

func NewWelcomeHandler() *WelcomeHandler {
	return &WelcomeHandler{}
}

func (handler *WelcomeHandler) Welcome(c *fiber.Ctx) error {
	return c.Render("welcome", nil, "layouts/main")
}
