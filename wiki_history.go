package backlog

import "time"

type WikiHistory struct {
	PageID      int64
	Version     int32
	Name        string
	Content     string
	CreatedUser User
	Created     time.Time
	UpdatedUser User
	Updated     time.Time
}
