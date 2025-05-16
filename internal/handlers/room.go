package handlers

import "github.com/gofiber/fiber/v2"

type RoomHandler struct{}

func NewRoomHandler() *RoomHandler {
	return &RoomHandler{}
}

func (handler *RoomHandler) CreateRoom(c *fiber.Ctx) error {
	return c.SendString("Create Room")
}

func (handler *RoomHandler) GetRoom(c *fiber.Ctx) error {
	return c.SendString("Get Room")
}

func (handler *RoomHandler) ChatRoom(c *fiber.Ctx) error {
	return c.SendString("Chat Room")
}

func (handler *RoomHandler) ChatRoomWS(c *fiber.Ctx) error {
	return c.SendString("Chat Room WS")
}

func (handler *RoomHandler) RoomViewerWS(c *fiber.Ctx) error {
	return c.SendString("Room Viewer WS")
}
