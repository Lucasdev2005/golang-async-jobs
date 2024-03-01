package types

import "time"

type User struct {
	UserID         int
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
	AccountLimit   int
	AccountBalance int
}
