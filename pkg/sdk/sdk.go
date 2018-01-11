package sdk

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"

	"github.com/jackspirou/syscerts"
	"golang.org/x/oauth2"
)

const (
	pathAuthLogin      = "%s/api/auth/login"
	pathProfile        = "%s/api/profile/self"
	pathProfileToken   = "%s/api/profile/token"
	pathUsers          = "%s/api/users"
	pathUser           = "%s/api/users/%v"
	pathUserTeam       = "%s/api/users/%v/teams"
	pathTeams          = "%s/api/teams"
	pathTeam           = "%s/api/teams/%v"
	pathTeamUser       = "%s/api/teams/%v/users"
)

// ClientAPI describes a client API.
type ClientAPI interface {
	// SetClient sets the default http client. This should
	// be used in conjunction with golang.org/x/oauth2 to
	// authenticate requests to the API.
	SetClient(client *http.Client)

	// IsAuthenticated checks if we already provided an authentication
	// token for our client requests. If it returns false you can update
	// the client after fetching a valid token.
	IsAuthenticated() bool

	// AuthLogin signs in based on credentials and returns a token.
	AuthLogin(string, string) (*Token, error)

	// ProfileToken returns a token.
	ProfileToken() (*Token, error)

	// ProfileGet returns a profile.
	ProfileGet() (*Profile, error)

	// ProfilePatch updates a profile.
	ProfilePatch(*Profile) (*Profile, error)

	// UserList returns a list of all users.
	UserList() ([]*User, error)

	// UserGet returns a user.
	UserGet(string) (*User, error)

	// UserPost creates a user.
	UserPost(*User) (*User, error)

	// UserPatch updates a user.
	UserPatch(*User) (*User, error)

	// UserDelete deletes a user.
	UserDelete(string) error

	// UserTeamList returns a list of related teams for a user.
	UserTeamList(UserTeamParams) ([]*TeamUser, error)

	// UserTeamAppend appends a team to a user.
	UserTeamAppend(UserTeamParams) error

	// UserTeamPerm updates perms for user team.
	UserTeamPerm(UserTeamParams) error

	// UserTeamDelete remove a team from a user.
	UserTeamDelete(UserTeamParams) error

	// TeamList returns a list of all teams.
	TeamList() ([]*Team, error)

	// TeamGet returns a team.
	TeamGet(string) (*Team, error)

	// TeamPost creates a team.
	TeamPost(*Team) (*Team, error)

	// TeamPatch updates a team.
	TeamPatch(*Team) (*Team, error)

	// TeamDelete deletes a team.
	TeamDelete(string) error

	// TeamUserList returns a list of related users for a team.
	TeamUserList(TeamUserParams) ([]*TeamUser, error)

	// TeamUserAppend appends a user to a team.
	TeamUserAppend(TeamUserParams) error

	// TeamUserPerm updates perms for team user.
	TeamUserPerm(TeamUserParams) error

	// TeamUserDelete remove a user from a team.
	TeamUserDelete(TeamUserParams) error
}

// Default implements the client interface.
type Default struct {
	client *http.Client
	base   string
	token  string
}

// NewClient returns a client for the specified URL.
func NewClient(uri string) ClientAPI {
	return &Default{
		client: http.DefaultClient,
		base:   uri,
	}
}

// NewClientToken returns a client that authenticates
// all outbound requests with the given token.
func NewClientToken(uri, token string) ClientAPI {
	config := oauth2.Config{}

	client := config.Client(
		context.Background(),
		&oauth2.Token{
			AccessToken: token,
		},
	)

	if trans, ok := client.Transport.(*oauth2.Transport); ok {
		trans.Base = &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			TLSClientConfig: &tls.Config{
				RootCAs: syscerts.SystemRootsPool(),
			},
		}
	}

	return &Default{
		client: client,
		base:   uri,
		token:  token,
	}
}

// IsAuthenticated checks if we already provided an authentication
// token for our client requests. If it returns false you can update
// the client after fetching a valid token.
func (c *Default) IsAuthenticated() bool {
	if c.token == "" {
		return false
	}

	uri, err := url.Parse(fmt.Sprintf(pathProfileToken, c.base))

	if err != nil {
		return false
	}

	req, err := http.NewRequest("GET", uri.String(), nil)

	if err != nil {
		return false
	}

	req.Header.Set(
		"User-Agent",
		UserAgent,
	)

	resp, err := c.client.Do(req)

	if err != nil {
		return false
	}

	defer resp.Body.Close()

	return resp.StatusCode == http.StatusUnauthorized
}

// SetClient sets the default http client. This should
// be used in conjunction with golang.org/x/oauth2 to
// authenticate requests to the API.
func (c *Default) SetClient(client *http.Client) {
	c.client = client
}

// AuthLogin signs in based on credentials and returns a token.
func (c *Default) AuthLogin(username, password string) (*Token, error) {
	out := &Token{}

	in := struct {
		Username string
		Password string
	}{
		username,
		password,
	}

	uri := fmt.Sprintf(pathAuthLogin, c.base)
	err := c.post(uri, in, out)

	return out, err
}

// ProfileToken returns a profile.
func (c *Default) ProfileToken() (*Token, error) {
	out := &Token{}

	uri := fmt.Sprintf(pathProfileToken, c.base)
	err := c.get(uri, out)

	return out, err
}

// ProfileGet returns a profile.
func (c *Default) ProfileGet() (*Profile, error) {
	out := &Profile{}

	uri := fmt.Sprintf(pathProfile, c.base)
	err := c.get(uri, out)

	return out, err
}

// ProfilePatch updates a profile.
func (c *Default) ProfilePatch(in *Profile) (*Profile, error) {
	out := &Profile{}

	uri := fmt.Sprintf(pathProfile, c.base)
	err := c.put(uri, in, out)

	return out, err
}

// UserList returns a list of all users.
func (c *Default) UserList() ([]*User, error) {
	var out []*User

	uri := fmt.Sprintf(pathUsers, c.base)
	err := c.get(uri, &out)

	return out, err
}

// UserGet returns a user.
func (c *Default) UserGet(id string) (*User, error) {
	out := &User{}

	uri := fmt.Sprintf(pathUser, c.base, id)
	err := c.get(uri, out)

	return out, err
}

// UserPost creates a user.
func (c *Default) UserPost(in *User) (*User, error) {
	out := &User{}

	uri := fmt.Sprintf(pathUsers, c.base)
	err := c.post(uri, in, out)

	return out, err
}

// UserPatch updates a user.
func (c *Default) UserPatch(in *User) (*User, error) {
	out := &User{}

	uri := fmt.Sprintf(pathUser, c.base, in.ID)
	err := c.put(uri, in, out)

	return out, err
}

// UserDelete deletes a user.
func (c *Default) UserDelete(id string) error {
	uri := fmt.Sprintf(pathUser, c.base, id)
	err := c.delete(uri, nil)

	return err
}

// UserTeamList returns a list of related teams for a user.
func (c *Default) UserTeamList(opts UserTeamParams) ([]*TeamUser, error) {
	var out []*TeamUser

	uri := fmt.Sprintf(pathUserTeam, c.base, opts.User)
	err := c.get(uri, &out)

	return out, err
}

// UserTeamAppend appends a team to a user.
func (c *Default) UserTeamAppend(opts UserTeamParams) error {
	uri := fmt.Sprintf(pathUserTeam, c.base, opts.User)
	err := c.post(uri, opts, nil)

	return err
}

// UserTeamPerm updates perms for user team.
func (c *Default) UserTeamPerm(opts UserTeamParams) error {
	uri := fmt.Sprintf(pathUserTeam, c.base, opts.User)
	err := c.put(uri, opts, nil)

	return err
}

// UserTeamDelete remove a team from a user.
func (c *Default) UserTeamDelete(opts UserTeamParams) error {
	uri := fmt.Sprintf(pathUserTeam, c.base, opts.User)
	err := c.delete(uri, opts)

	return err
}

// TeamList returns a list of all teams.
func (c *Default) TeamList() ([]*Team, error) {
	var out []*Team

	uri := fmt.Sprintf(pathTeams, c.base)
	err := c.get(uri, &out)

	return out, err
}

// TeamGet returns a team.
func (c *Default) TeamGet(id string) (*Team, error) {
	out := &Team{}

	uri := fmt.Sprintf(pathTeam, c.base, id)
	err := c.get(uri, out)

	return out, err
}

// TeamPost creates a team.
func (c *Default) TeamPost(in *Team) (*Team, error) {
	out := &Team{}

	uri := fmt.Sprintf(pathTeams, c.base)
	err := c.post(uri, in, out)

	return out, err
}

// TeamPatch updates a team.
func (c *Default) TeamPatch(in *Team) (*Team, error) {
	out := &Team{}

	uri := fmt.Sprintf(pathTeam, c.base, in.ID)
	err := c.put(uri, in, out)

	return out, err
}

// TeamDelete deletes a team.
func (c *Default) TeamDelete(id string) error {
	uri := fmt.Sprintf(pathTeam, c.base, id)
	err := c.delete(uri, nil)

	return err
}

// TeamUserList returns a list of related users for a team.
func (c *Default) TeamUserList(opts TeamUserParams) ([]*TeamUser, error) {
	var out []*TeamUser

	uri := fmt.Sprintf(pathTeamUser, c.base, opts.Team)
	err := c.get(uri, &out)

	return out, err
}

// TeamUserAppend appends a user to a team.
func (c *Default) TeamUserAppend(opts TeamUserParams) error {
	uri := fmt.Sprintf(pathTeamUser, c.base, opts.Team)
	err := c.post(uri, opts, nil)

	return err
}

// TeamUserPerm updates perms for team user.
func (c *Default) TeamUserPerm(opts TeamUserParams) error {
	uri := fmt.Sprintf(pathTeamUser, c.base, opts.Team)
	err := c.put(uri, opts, nil)

	return err
}

// TeamUserDelete remove a user from a team.
func (c *Default) TeamUserDelete(opts TeamUserParams) error {
	uri := fmt.Sprintf(pathTeamUser, c.base, opts.Team)
	err := c.delete(uri, opts)

	return err
}
