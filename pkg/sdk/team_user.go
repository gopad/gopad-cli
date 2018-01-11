package sdk

// TeamUser represents a team user API response.
type TeamUser struct {
	Team *Team  `json:"team,omitempty"`
	User *User  `json:"user,omitempty"`
	Perm string `json:"perm,omitempty"`
}
