package server

import (
	"flag"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"
	"moraes-streaming/internal/handlers"
	w "moraes-streaming/pkg/webrtc"
	"time"
)

var (
	addr = flag.String("addr", ":8080", "http service address")
	cert = flag.String("cert", "cert.pem", "https certificate")
	key  = flag.String("key", "key.pem", "https key")
)

func Run() error {
	flag.Parse()

	if *addr == ":" {
		*addr = ":8080"
	}

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(cors.New())

	// handlers
	chatHandler := handlers.NewChatHandler()
	_ = chatHandler
	roomHandler := handlers.NewRoomHandler()
	streamHandler := handlers.NewStreamHandler()
	welcomeHandler := handlers.NewWelcomeHandler()

	app.Get("/", welcomeHandler.Welcome)

	app.Get("/room/create", roomHandler.CreateRoom)

	app.Get("/room/:id", roomHandler.GetRoom)

	app.Get("/room/:id/ws", websocket.New(roomHandler.RoomWS, websocket.Config{
		HandshakeTimeout: 10 * time.Second,
	}))

	app.Get("/room/:id/chat", roomHandler.RoomChat)

	app.Get("/room/:id/chat/ws", websocket.New(roomHandler.RoomChatWS, websocket.Config{
		HandshakeTimeout: 10 * time.Second,
	}))

	app.Get("/room/:id/viewer/ws", websocket.New(roomHandler.RoomViewerWS, websocket.Config{
		HandshakeTimeout: 10 * time.Second,
	}))

	app.Get("/stream/:id", streamHandler.Stream)

	app.Get("/stream/:id/ws", websocket.New(streamHandler.StreamWS, websocket.Config{
		HandshakeTimeout: 10 * time.Second,
	}))

	app.Get("/stream/:id/chat/ws", websocket.New(streamHandler.StreamChatWS, websocket.Config{
		HandshakeTimeout: 10 * time.Second,
	}))

	app.Get("/stream/:id/viewer/ws", websocket.New(streamHandler.StreamViewerWS, websocket.Config{
		HandshakeTimeout: 10 * time.Second,
	}))

	app.Static("/", "./assets")

	w.Rooms = make(map[string]*w.Room)
	w.Streams = make(map[string]*w.Room)

	go dispatchKeyFrames()

	if *cert != "" {
		return app.ListenTLS(*addr, *cert, *key)
	}

	return app.Listen(*addr)
}

func dispatchKeyFrames() {
	for range time.NewTicker(3 * time.Second).C {
		for _, room := range w.Rooms {
			room.Peers.DispatchKeyFrame()
		}
	}
}
