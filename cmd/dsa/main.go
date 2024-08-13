package main

import (
	"BryanMcGahan/dsa_viz/internal/dsa/handlers"
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

const (
	PORT        = "80"
	NUM_OF_NUMS = 10
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {

	app := fiber.New()

	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws/hello", websocket.New(handlers.BubbleSort))
	log.Fatal(app.Listen(":" + PORT))
}
