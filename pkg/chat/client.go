package chat

import (
	"github.com/fasthttp/websocket"
	"time"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingWait       = (pongWait * 9) / 10
	maxMessageSize = 512
)

type Client struct {
	Hub  *Hub
	Conn *websocket.Conn
	Send chan []byte
}

func (c *Client) readPump() {

}

func (c *Client) writePump() {

}

func PeerChatConn() {

}
