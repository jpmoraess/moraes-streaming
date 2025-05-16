package server

import (
	"flag"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"
	"moraes-streaming/internal/handlers"
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
	app.Get("/room/:id/chat", roomHandler.ChatRoom)
	app.Get("/room/:id/chat/ws", roomHandler.ChatRoomWS)
	app.Get("/room/:id/viewer/ws", roomHandler.RoomViewerWS)
	app.Get("/stream/:id", streamHandler.Stream)
	app.Get("/stream/:id/ws", streamHandler.StreamWS)
	app.Get("/stream/:id/chat/ws")
	app.Get("/stream/:id/viewer/ws")

	return app.Listen(*addr)
}
