package rest

import (
	"bSocial/domain"
	"bSocial/helpers"
	"bSocial/interface/kafkaProducer"
	"bSocial/interface/mysql"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func InitCommentApi(api fiber.Router) {
	api.Post("/comment/:postID", apiPostComment())
	api.Get("/comment/:postID", apiGetComment())
	api.Get("/notifications", apiGetNewNotification())
}

// creates a user comment for specific postID
func apiPostComment() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		getIntPostID, err := strconv.ParseInt(c.Params("postID"), 10, 32)
		if err != nil {
			return ResponseWithError(c, err.Error(), nil)
		}
		commentData := new(domain.Comment)
		if err := c.BodyParser(commentData); err != nil {
			return ResponseWithError(c, err.Error(), nil)
		}
		if errors := domain.ValidateType(*commentData); errors != nil {
			return ResponseWithError(c, "", errors)
		}
		commentData.PostID = uint(getIntPostID)
		newComment, err := mysql.CreateComment(commentData)

		if err != nil {
			return ResponseWithError(c, err.Error(), nil)
		}

		// TODO: handle error
		kafkaComment, err := mysql.GetCommentForKafka(newComment.ID)

		if err != nil {
			return ResponseWithError(c, err.Error(), nil)
		}

		kafkaProducer.ProduceComment(*kafkaComment)

		return ResponseWithData(c, newComment)
	}
}

func apiGetComment() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		getIntPostID, err := strconv.ParseInt(c.Params("postID"), 10, 32)
		if err != nil {
			return ResponseWithError(c, err.Error(), nil)
		}

		comments, err := mysql.GetCommentsForPost(uint(getIntPostID))
		if err != nil {
			return ResponseWithError(c, err.Error(), nil)
		}
		return ResponseWithData(c, comments)
	}
}

func apiGetNewNotification() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		thisUserID := helpers.ExtractJWTUserID(c)

		comments, err := mysql.GetUndeliveredComments(thisUserID)
		if err != nil {
			return ResponseWithError(c, err.Error(), nil)
		}
		return ResponseWithData(c, comments)
	}
}
