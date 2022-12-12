package mysql

import (
	"bSocial/domain"
)

func CreateComment(commentData *domain.Comment) (*domain.Comment, error) {
	var post domain.Comment

	if resultCreate := MySql.Create(commentData); resultCreate.Error != nil {
		return nil, GenericDBError()
	}

	// need to fetch comment again because date have no CreatedAt value after create
	if resultSelect := MySql.Where("id = ?", commentData.ID).First(&post); resultSelect.Error != nil {
		return nil, GenericDBError()
	}

	return &post, nil
}

func GetCommentsForPost(postID uint) ([]domain.Comment, error) {
	var comments []domain.Comment
	if result := MySql.Where("post_id = ?", postID).Find(&comments); result.Error != nil {
		return nil, GenericDBError()
	}
	return comments, nil
}

func GetCommentForKafka(commentID uint) (*domain.KafkaComment, error) {
	var kafkaComment domain.KafkaComment
	result := MySql.Table("comments").
		Where("comments.id = ?", commentID).
		Select("comments.id, comments.created_at, comments.content, comments.post_id, posts.user_id, users.username, users.email").
		Joins("join posts on comments.post_id = posts.id").
		Joins("join users on posts.user_id = users.id").
		Scan(&kafkaComment)
	if result.Error != nil {
		return nil, GenericDBError()
	}
	return &kafkaComment, nil
}

func GetUndeliveredComments(userID uint) ([]domain.Comment, error) {
	var comments []domain.Comment
	var commentIDS []uint
	if result := MySql.
		Where("posts.user_id = ?", userID).
		Where("comments.delivered = ?", false).
		Joins("join posts on comments.post_id = posts.id").
		Find(&comments); result.Error != nil {

		return nil, GenericDBError()
	}
	for _, comment := range comments {
		commentIDS = append(commentIDS, comment.ID)
	}
	if resultUpdate := MySql.Table("comments").Where("id IN ?", commentIDS).Update("delivered", true); resultUpdate.Error != nil {
		return nil, GenericDBError()
	}

	return comments, nil
}
