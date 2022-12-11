package main

import (
	"bSocial/helpers"
	"bSocial/interface/mysql"
	"bSocial/interface/rest"
	"log"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func main() {
	app := fiber.New()
	err := helpers.InitConfig()
	if err != nil {
		log.Fatalf("Error initialising config: %s", err)
	}
	public := app.Group("/api")
	rest.InitAuthApi(public)

	private := app.Group("/api")
	private.Use(jwtware.New(jwtware.Config{SigningKey: []byte(helpers.CONFIG.Json.Secret)}))
	rest.InitUserApi(private)
	rest.InitPostApi(private)
	rest.InitCommentApi(private)

	err = mysql.InitConnection()
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}

	mysql.AutoMigrate()

	app.Listen(":5000")
}
