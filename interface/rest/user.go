package rest

import (
	"bSocial/interface/mysql"

	"github.com/gofiber/fiber/v2"
)

func InitUserApi(api fiber.Router) {
	api.Get("/user", apiGetUsers())
}

func apiGetUsers() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		users, _ := mysql.GetUsers()
		return c.JSON(users)
	}
}
