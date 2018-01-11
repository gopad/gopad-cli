package sdk

// Token represents a session token.
type Token struct {
	Token  string `json:"token"`
	Expire string `json:"expire,omitempty"`
}
