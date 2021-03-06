package entity

import (
	"time"
)

// User holds the user entity fields
type User struct {
	ID           int
	FirstName    string `db:"first_name"`
	LastName     string `db:"last_name"`
	MobileNumber string `db:"mobile_number"`
	Metadata     string
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
	Password     string
}

// GetModelName returns the model name of academic year entity that can be used for naming schemas
func (entity *User) GetModelName() string {
	return "users"
}
