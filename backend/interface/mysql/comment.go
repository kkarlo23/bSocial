package mysql

import (
	"bSocial/domain"
)

func CreateComment(commentData *domain.Comment) (*domain.Comment, error) {
	var post domain.Comment

	if resultCreate := MySql.Create(commentData); resultCreate.Error != nil {
		return nil, resultCreate.Error
	}

	// need to fetch user again because date have no value after create
	if resultSelect := MySql.Where("id = ?", commentData.ID).First(&post); resultSelect.Error != nil {
		return nil, resultSelect.Error
	}

	return &post, nil
}
