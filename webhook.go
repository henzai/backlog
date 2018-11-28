package backlog

import (
	"time"
)

type Webhook struct {
	ID              int64
	Name            string
	Description     string
	HookURL         string
	AllEvent        bool
	ActivityTypeIds []int
	CreatedUser     User
	Created         time.Time
	UpdatedUser     User
	Updated         time.Time
}
