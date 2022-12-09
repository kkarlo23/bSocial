package mysql

import (
	"bSocial/domain"
)

func CreatePost(postData *domain.Post) (*domain.Post, error) {
	var post domain.Post

	if resultCreate := MySql.Create(postData); resultCreate.Error != nil {
		return nil, GenericDBError()
	}

	// need to fetch post again because CreatedAt have no value after create
	if resultSelect := MySql.Where("id = ?", postData.ID).First(&post); resultSelect.Error != nil {
		return nil, GenericDBError()
	}

	return &post, nil
}

func GetPostsForUser(userID uint) ([]domain.Post, error) {
	var followingIDS []uint
	var posts []domain.Post
	result := MySql.Table("user_followers").Where("follower_id = ?", userID).Select("following_id").Find(&followingIDS)
	if result.Error != nil {
		return nil, GenericDBError()
	}
	result = MySql.Where("user_id IN ?", followingIDS).Find(&posts)
	if result.Error != nil {
		return nil, GenericDBError()
	}
	return posts, nil
}

func GetPostsForKafka(postID uint) (*domain.KafkaPost, error) {
	var kafkaPost domain.KafkaPost
	result := MySql.Table("posts").
		Where("posts.id = ?", postID).
		Select("posts.id, posts.created_at, posts.content, posts.user_id, users.username, users.email").
		Joins("join users on posts.user_id = users.id").
		Scan(&kafkaPost)
	if result.Error != nil {
		return nil, GenericDBError()
	}
	return &kafkaPost, nil
}
