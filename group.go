package backlog

import (
	"time"
)

type Group struct {
	ID           int64
	Name         string
	Members      []User
	DisplayOrder int64
	CreatedUser  User
	Created      time.Time
	UpdatedUser  User
	Updated      time.Time
}
