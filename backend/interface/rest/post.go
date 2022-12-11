package rest

import (
	"bSocial/domain"
	"bSocial/helpers"
	"bSocial/interface/kafkaProducer"
	"bSocial/interface/mysql"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func InitPostApi(api fiber.Router) {
	api.Post("/post", apiPostPost())
	api.Get("/post", apiGetPosts())
}

// creates a posts for current (logged in) user
func apiPostPost() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		thisUserID := helpers.ExtractJWTUserID(c)
		postData := new(domain.Post)
		if err := c.BodyParser(postData); err != nil {
			return ResponseWithError(c, err.Error(), nil)
		}
		if errors := domain.ValidateType(*postData); errors != nil {
			return ResponseWithError(c, "", errors)
		}
		postData.UserID = thisUserID
		newPost, err := mysql.CreatePost(postData)
		if err != nil {
			return ResponseWithError(c, err.Error(), nil)
		}

		kafkaPost, err := mysql.GetPostsForKafka(newPost.ID)

		if err != nil {
			return ResponseWithError(c, err.Error(), nil)
		}

		// TODO: handle error
		kafkaProducer.ProducePost(*kafkaPost)

		return ResponseWithData(c, newPost)
	}
}

// returns all posts from users that current (logged in) user is following
func apiGetPosts() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		thisUserID := helpers.ExtractJWTUserID(c)

		rowsPerPage, err := strconv.ParseInt(c.Query("rowsPerPage"), 10, 64)
		if err != nil {
			rowsPerPage = 10
		}

		page, err := strconv.ParseInt(c.Query("page"), 10, 64)
		if err != nil {
			page = 1
		}

		posts, err := mysql.GetPostsForUser(thisUserID, rowsPerPage, page)
		if err != nil {
			return ResponseWithError(c, err.Error(), nil)
		}

		return ResponseWithData(c, posts)
	}
}
