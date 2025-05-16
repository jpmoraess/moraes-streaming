package handlers

import (
	"fmt"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type RoomHandler struct{}

func NewRoomHandler() *RoomHandler {
	return &RoomHandler{}
}

func (handler *RoomHandler) CreateRoom(c *fiber.Ctx) error {
	return c.Redirect(fmt.Sprintf("/room/%s", uuid.New().String()))
}

func (handler *RoomHandler) RoomWS(c *websocket.Conn) {
	roomId := c.Params("id")
	if roomId == "" {
		return
	}

	createOrGetRoom(uuid.MustParse(roomId))
}

func (handler *RoomHandler) GetRoom(c *fiber.Ctx) error {
	roomId := c.Params("id")
	if roomId == "" {
		return c.SendStatus(fiber.StatusNotFound)
	}

	createOrGetRoom(uuid.MustParse(roomId))
	return nil
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

func createOrGetRoom(id uuid.UUID) (string, string, string) {
	return "", "", ""
}
