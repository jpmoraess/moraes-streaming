package handlers

import (
	"fmt"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	w "moraes-streaming/pkg/webrtc"
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

func (handler *RoomHandler) RoomChat(c *fiber.Ctx) error {
	return c.SendString("Room Chat")
}

func (handler *RoomHandler) RoomChatWS(c *websocket.Conn) {
	roomId := c.Params("id")
	if roomId == "" {
		return
	}

	createOrGetRoom(uuid.MustParse(roomId))
}

func (handler *RoomHandler) RoomViewerWS(c *websocket.Conn) {
	roomId := c.Params("id")
	if roomId == "" {
		return
	}

	createOrGetRoom(uuid.MustParse(roomId))
}

func createOrGetRoom(id uuid.UUID) (string, string, *w.Room) {
	return "", "", ""
}

func roomViewerWS(c *websocket.Conn) {

}

func roomViewerConn(c *websocket.Conn, p *w.Peers) {

}

type Message struct {
	Event string `json:"event"`
	Data  string `json:"data"`
}
