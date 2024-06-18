package model

import "time"

type Sneaker struct {
	ID        string
	Title     string
	Price     int
	ImageUrl  string
	Available bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
