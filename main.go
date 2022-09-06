package main

import (
	"fiber_kafka/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	router.Router(app)
	app.Listen(":3000")
}
