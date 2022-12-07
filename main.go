package main

import (
	"bSocial/helpers"
	"bSocial/interface/mysql"
	"bSocial/interface/rest"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func main() {
	app := fiber.New()

	public := app.Group("/api")
	rest.InitAuthApi(public)

	private := app.Group("/api")
	private.Use(jwtware.New(jwtware.Config{SigningKey: []byte("secret")}))
	rest.InitUserApi(private)

	args := helpers.GetArgs()
	mysql.InitConnection(args)
	mysql.AutoMigrate()

	app.Listen(":5000")
}
