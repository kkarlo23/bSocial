package domain

import "time"

type KafkaComment struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Content   string    `json:"content"`
	PostID    string    `json:"postID"`
	UserID    uint      `json:"userID"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
}
