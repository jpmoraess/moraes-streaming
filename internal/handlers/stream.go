package handlers

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type StreamHandler struct{}

func NewStreamHandler() *StreamHandler {
	return &StreamHandler{}
}

func (handler *StreamHandler) Stream(c *fiber.Ctx) error {
	return c.SendString("Stream....")
}

func (handler *StreamHandler) StreamWS(c *websocket.Conn) {
	streamId := c.Params("id")
	if streamId == "" {
		return
	}
}

func (handler *StreamHandler) StreamChatWS(c *websocket.Conn) {
	streamId := c.Params("id")
	if streamId == "" {
		return
	}
}

func (handler *StreamHandler) StreamViewerWS(c *websocket.Conn) {
	streamId := c.Params("id")
	if streamId == "" {
		return
	}
}
