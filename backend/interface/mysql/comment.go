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
