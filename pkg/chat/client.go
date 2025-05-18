package chat

import (
	"bytes"
	"fmt"
	"github.com/fasthttp/websocket"
	"log"
	"time"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingWait       = (pongWait * 9) / 10
	maxMessageSize = 512
)

var (
	newLine = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.FastHTTPUpgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Client struct {
	Hub  *Hub
	Conn *websocket.Conn
	Send chan []byte
}

func (c *Client) readPump() {
	defer func() {
		c.Hub.unregister <- c
		if err := c.Conn.Close(); err != nil {
			fmt.Println("close connection error:", err)
		}
	}()
	c.Conn.SetReadLimit(maxMessageSize)

	if err := c.Conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		fmt.Println("set read deadline error:", err)
	}

	c.Conn.SetPongHandler(func(string) error {
		if err := c.Conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
			fmt.Println("set read deadline error:", err)
			return err
		}
		return nil
	})

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newLine, space, -1))
		c.Hub.broadcast <- message
	}
}

func (c *Client) writePump() {

}

func PeerChatConn(c *websocket.Conn, hub *Hub) {
	client := &Client{
		Hub:  hub,
		Conn: c,
		Send: make(chan []byte, 256),
	}
	client.Hub.register <- client

	go client.writePump()
	client.readPump()
}
