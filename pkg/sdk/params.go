package sdk

// TeamUserParams is used to assign users to a team.
type TeamUserParams struct {
	Team string `json:"team"`
	User string `json:"user"`
	Perm string `json:"perm"`
}

// UserTeamParams is used to assign teams to a user.
type UserTeamParams struct {
	User string `json:"user"`
	Team string `json:"team"`
	Perm string `json:"perm"`
}
