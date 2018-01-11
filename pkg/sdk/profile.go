package sdk

import (
	"time"
)

// Profile represents a profile API response.
type Profile struct {
	ID        int64     `json:"id"`
	Slug      string    `json:"slug"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Active    bool      `json:"active"`
	Admin     bool      `json:"admin"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s *Profile) String() string {
	return s.Username
}
