package rest

import (
	"bSocial/domain"
	"bSocial/helpers"
	"bSocial/interface/mysql"

	"github.com/gofiber/fiber/v2"
)

func InitPostApi(api fiber.Router) {
	api.Post("/post", apiPostPost())
	api.Get("/post", apiGetPosts())
}

func apiPostPost() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		thisUserID := helpers.ExtractJWTUserID(c)
		postData := new(domain.Post)
		if err := c.BodyParser(postData); err != nil {
			return err
		}
		if errors := domain.ValidateType(*postData); errors != nil {
			return c.Status(fiber.StatusBadRequest).JSON(errors)
		}
		postData.UserID = thisUserID
		newPost, _ := mysql.CreatePost(postData)

		return c.JSON(newPost)
	}
}

func apiGetPosts() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		thisUserID := helpers.ExtractJWTUserID(c)

		posts, _ := mysql.GetPostsForUser(thisUserID)

		return c.JSON(posts)
	}
}
