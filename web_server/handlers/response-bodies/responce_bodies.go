package responsebodies

import "time"

type SessionToken struct {
	Token  string    `json:"session_token"`
	Expiry time.Time `json:"expires_at"`
}
