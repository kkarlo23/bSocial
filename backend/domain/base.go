package domain

import "time"

// Base is an struct that all other entities can use
// there could be more fields, for example UpdatedAt, DeletedAt(if you do not want to actualy delete users from database)
type Base struct {
	CreatedAt time.Time `json:"createdAt" gorm:"type:datetime;default:NOW();not null;"`
}
