package domain

import "time"

type Post struct {
	ID       uint      `json:"id" gorm:"primaryKey;autoIncrement;not null;"`
	Content  string    `json:"content" gorm:"type:varchar(1024);not null" validate:"required,min=3,max=1024"`
	UserID   uint      `json:"userID"`
	Comments []Comment `json:"-" gorm:"foreignKey:PostID"`
	Base
}

type KafkaPost struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Content   string    `json:"content"`
	UserID    uint      `json:"userID"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
}
