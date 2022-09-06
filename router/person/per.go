package person

import "github.com/gofiber/fiber/v2"

func Person_request(c *fiber.Ctx) error {
	return c.SendString("this person ctx")
}
