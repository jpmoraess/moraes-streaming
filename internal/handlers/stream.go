package handlers

import "github.com/gofiber/fiber/v2"

type StreamHandler struct{}

func NewStreamHandler() *StreamHandler {
	return &StreamHandler{}
}

func (handler *StreamHandler) Stream(c *fiber.Ctx) error {
	return c.SendString("Stream....")
}

func (handler *StreamHandler) StreamWS(c *fiber.Ctx) error {
	return c.SendString("Stream WS")
}
