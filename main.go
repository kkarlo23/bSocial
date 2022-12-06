package main

import (
	"bSocial/helpers"
	"bSocial/interface/mysql"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	args := helpers.GetArgs()
	mysql.InitConnection(args)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Listen(":5000")
}
