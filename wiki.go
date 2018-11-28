package backlog

import "time"

type Wiki struct {
	ID          int64
	projectId   int64
	Name        string
	Content     string
	Tags        []WikiTag
	Attachments []Attachment
	SharedFiles []SharedFile
	Stars       []Star
	CreatedUser User
	Created     time.Time
	UpdatedUser User
	Updated     time.Time
}

type WikiTag struct {
	ID   int64
	Name string
}
