package mysql

import (
	"bSocial/domain"
)

func CreatePost(postData *domain.Post) (*domain.Post, error) {
	var post domain.Post

	if resultCreate := MySql.Create(postData); resultCreate.Error != nil {
		return nil, resultCreate.Error
	}

	// need to fetch user again because date have no value after create
	if resultSelect := MySql.Where("id = ?", postData.ID).First(&post); resultSelect.Error != nil {
		return nil, resultSelect.Error
	}

	return &post, nil
}

func GetPostsForUser(userID uint) ([]domain.Post, error) {
	var followingIDS []uint
	var posts []domain.Post
	result := MySql.Table("user_followers").Where("follower_id = ?", userID).Select("following_id").Find(&followingIDS)
	if result.Error != nil {
		return nil, result.Error
	}
	result = MySql.Where("user_id IN ?", followingIDS).Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	return posts, nil
}
