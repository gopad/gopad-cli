package sdk

import (
	"time"
)

// User represents a user API response.
type User struct {
	ID        int64     `json:"id"`
	Slug      string    `json:"slug"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Active    bool      `json:"active"`
	Admin     bool      `json:"admin"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Teams     []*Team   `json:"teams,omitempty"`
}

func (s *User) String() string {
	return s.Username
}
