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
	err := helpers.InitConfig()
	if err != nil {
		panic(err)
	}
	public := app.Group("/api")
	rest.InitAuthApi(public)

	private := app.Group("/api")
	private.Use(jwtware.New(jwtware.Config{SigningKey: []byte(helpers.CONFIG.Json.Secret)}))
	rest.InitUserApi(private)
	rest.InitPostApi(private)
	rest.InitCommentApi(private)

	mysql.InitConnection()
	mysql.AutoMigrate()

	app.Listen(":5000")
}
