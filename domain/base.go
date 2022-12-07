package domain

import "time"

// Base is an struct that all other entities will contain
type Base struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement;not null;"`
	CreatedAt time.Time `json:"createdAt" gorm:"type:datetime;default:NOW();not null;"`
}
