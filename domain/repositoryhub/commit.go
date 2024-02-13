package repositoryhub

import "time"

type Commit struct {
	SHA         string
	Name        string
	Description string
	Date        time.Time
}
