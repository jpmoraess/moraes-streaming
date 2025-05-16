package handlers

import "github.com/gofiber/fiber/v2"

type ChatHandler struct{}

func NewChatHandler() *ChatHandler {
	return &ChatHandler{}
}

func (handler *ChatHandler) HandleChat(c *fiber.Ctx) error {
	return c.SendString("Chatting...")
}
