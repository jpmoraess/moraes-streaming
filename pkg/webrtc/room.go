package webrtc

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/pion/webrtc/v3"
	"log"
	"moraes-streaming/pkg/chat"
)

type Room struct {
	Peers *Peers
	Hub   *chat.Hub
}

func RoomConn(c *websocket.Conn, p *Peers) {
	var config webrtc.Configuration

	peerConnection, err := webrtc.NewPeerConnection(config)
	if err != nil {
		log.Print(err)
		return
	}
	_ = peerConnection
}
