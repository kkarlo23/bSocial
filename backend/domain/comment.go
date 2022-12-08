package domain

type Comment struct {
	ID      uint   `json:"id" gorm:"primaryKey;autoIncrement;not null;"`
	Content string `json:"content" gorm:"type:varchar(1024);not null" validate:"required,min=3,max=1024"`
	PostID  uint   `json:"postID"`
	Base
}