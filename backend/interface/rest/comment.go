package rest

import (
	"bSocial/domain"
	"bSocial/interface/mysql"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func InitCommentApi(api fiber.Router) {
	api.Post("/comment/:postID", apiPostComment())
}

func apiPostComment() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		getIntPostID, err := strconv.ParseInt(c.Params("postID"), 10, 32)
		if err != nil {
			return err
		}
		commentData := new(domain.Comment)
		if err := c.BodyParser(commentData); err != nil {
			return err
		}
		if errors := domain.ValidateType(*commentData); errors != nil {
			return c.Status(fiber.StatusBadRequest).JSON(errors)
		}
		commentData.PostID = uint(getIntPostID)
		newComment, _ := mysql.CreateComment(commentData)

		return c.JSON(newComment)
	}
}
