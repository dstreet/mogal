package user

import "time"

type User struct {
	ID        string
	Email     string
	Active    bool
	CreatedAt time.Time
	DeletedAt *time.Time
	LastLogin time.Time
}
