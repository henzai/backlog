package backlog

import "time"

type IssueComment struct {
	ID            int64
	Content       string
	ChangeLog     []ChangeLog
	CreatedUser   User
	Created       time.Time
	Updated       time.Time
	Stars         []Star
	Notifications []Notification
}

type ChangeLog struct {
	Field            string
	NewValue         string
	OriginalValue    string
	AttachmentInfo   AttachmentInfo
	AttributeInfo    AttributeInfo
	NotificationInfo NotificationInfo
}

type AttachmentInfo struct {
	ID   int64
	Name string
}

type AttributeInfo struct {
	ID     int64
	TypeID string
}

type NotificationInfo struct {
	Type string
}

type Notification struct {
	ID                  int64
	AlreadyRead         bool
	Reason              int
	resourceAlreadyRead bool
	Project             Project
	Issue               Issue
	Comment             IssueComment
	PullRequest         interface{}
	PullRequestComment  interface{}
	Sender              User
	User                User
	Created             time.Time
}
