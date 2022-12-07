package domain

type User struct {
	Base
	FirstName string `json:"firstName" gorm:"type:varchar(100);not null" validate:"required,min=3,max=100"`
	LastName  string `json:"lastName" gorm:"type:varchar(100);not null" validate:"required,min=3,max=100"`
	Username  string `json:"username" gorm:"type:varchar(100);not null;unique" validate:"required,min=3,max=100"`
	Email     string `json:"email" gorm:"type:varchar(100);not null;unique" validate:"required,min=3,max=100"`
	Password  string `json:"-" gorm:"type:varbinary(100);not null;"` //->:false;<-:create
	// password is omited in json, there is no need to show it anywhere, also, gorm is only able to write to password, read is disabled
}
