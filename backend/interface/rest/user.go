package rest

import (
	"bSocial/helpers"
	"bSocial/interface/mysql"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func InitUserApi(api fiber.Router) {
	api.Get("/user", apiGetUsers())
	api.Post("/user/follow/:userID", apiPostUserFollow())
}

func apiGetUsers() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		users, _ := mysql.GetUsers()
		return c.JSON(users)
	}
}

func apiPostUserFollow() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		thisUserID := helpers.ExtractJWTUserID(c)
		getUserIDInt64, err := strconv.ParseInt(c.Params("userID"), 10, 32)
		if err != nil {
			return err
		}
		getUserID := uint(getUserIDInt64)

		err = mysql.UserFollow(thisUserID, getUserID)
		if err != nil {
			return err
		}
		return c.JSON("created")
	}
}
