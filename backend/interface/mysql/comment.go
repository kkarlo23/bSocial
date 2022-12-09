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
