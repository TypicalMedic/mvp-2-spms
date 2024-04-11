package session

import "time"

// this map stores the users sessions. For larger scale applications, you can use a database or cache for this purpose
var Sessions = map[string]Session{}

const SessionDefaultExpTime = 7 * 24 * time.Hour

type Session struct {
	user   userInfo
	expiry time.Time
}

func (s Session) IsExpired() bool {
	return s.expiry.Before(time.Now())
}

func (s Session) GetUser() userInfo {
	return s.user
}

func InitSession(user userInfo, exp time.Time) Session {
	return Session{
		user:   user,
		expiry: exp,
	}
}

type userInfo struct {
	username    string
	professorId string
}

func InitUserInfo(username string, profId string) userInfo {
	return userInfo{
		username:    username,
		professorId: profId,
	}
}

func (u userInfo) GetUsername() string {
	return u.username
}
func (u userInfo) GetProfId() string {
	return u.professorId
}
