package user

import "time"

type User struct {
	ID         int
	Name       string
	Email      string
	Occupation string
	Password   string
	Avatar     string
	Role       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
