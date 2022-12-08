package domain

import "time"

// Base is an struct that all other entities will contain
type Base struct {
	CreatedAt time.Time `json:"createdAt" gorm:"type:datetime;default:NOW();not null;"`
}
