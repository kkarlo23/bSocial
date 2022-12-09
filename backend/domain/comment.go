package domain

import "time"

type Comment struct {
	ID      uint   `json:"id" gorm:"primaryKey;autoIncrement;not null;"`
	Content string `json:"content" gorm:"type:varchar(1024);not null" validate:"required,min=3,max=1024"`
	PostID  uint   `json:"postID"`
	Base
}

type KafkaComment struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Content   string    `json:"content"`
	PostID    string    `json:"postID"`
	UserID    uint      `json:"userID"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
}
