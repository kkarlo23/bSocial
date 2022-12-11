package domain

import "time"

type Base struct {
	CreatedAt time.Time `json:"createdAt"`
}
