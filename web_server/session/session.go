package session

import "time"

type Session struct {
	username string
	expiry   time.Time
}

func (s Session) IsExpired() bool {
	return s.expiry.Before(time.Now())
}

func (s Session) GetUsername() string {
	return s.username
}

func InitSession(username string, exp time.Time) Session {
	return Session{
		username: username,
		expiry:   exp,
	}
}

// this map stores the users sessions. For larger scale applications, you can use a database or cache for this purpose
var Sessions = map[string]Session{}
