package domain

type User struct {
	ID        uint            `json:"id" gorm:"primaryKey;autoIncrement;not null;"`
	FirstName string          `json:"firstName" gorm:"type:varchar(100);not null" validate:"required,min=3,max=100"`
	LastName  string          `json:"lastName" gorm:"type:varchar(100);not null" validate:"required,min=3,max=100"`
	Username  string          `json:"username" gorm:"type:varchar(100);not null;unique" validate:"required,min=3,max=100"`
	Email     string          `json:"email" gorm:"type:varchar(100);not null;unique" validate:"required,min=3,max=100"`
	Password  string          `json:"-" gorm:"type:varbinary(100);not null;"`
	Posts     []Post          `json:"-" gorm:"foreignKey:UserID"`
	Following []*UserFollower `json:"-" gorm:"foreignKey:FollowerID;references:ID"`
	Followers []*UserFollower `json:"-" gorm:"foreignKey:FollowingID;references:ID"`
	Base
	// password is omited in json, there is no need to show it anywhere
}

type UserFollower struct {
	Base
	FollowerID  uint `gorm:"primaryKey;autoIncrement:false;not null"`
	FollowingID uint `gorm:"primaryKey;autoIncrement:false;not null"`
}
