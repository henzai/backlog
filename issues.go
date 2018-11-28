package backlog

import (
	"fmt"
	"strconv"
	"time"
)

type Issue struct {
	ID             int
	ProjectID      int
	IssueKey       string
	KeyID          int
	IssueType      IssueType
	Summary        string
	Description    string
	Resolution     Resolution
	Priority       Priority
	Status         Status
	Assignee       User
	Category       []Category
	Versions       []Version
	MileStone      []MileStone
	StartDate      time.Time
	DueDate        time.Time
	EstimatedHours float32
	ActualHours    float32
	ParentIssueID  int
	CreatedUser    User
	Created        time.Time
	UpdatedUser    User
	Updated        time.Time
	CustomFields   []CustomField
	Attachments    []Attachment
	SharedFiles    []SharedFile
	Stars          []Star
}

type Resolution struct {
	ID   int32
	Name string
}

type SharedFile struct {
	ID          int64
	Type        string
	Name        string
	Dir         string
	size        int64
	CreatedUser User
	Created     time.Time
	UpdatedUser User
	Updated     time.Time
}

type IssueType struct {
	ID        int64
	ProjectID int64
	Name      string
	Color     string
}

type Priority struct {
	ID   int
	Name string
}

type Status struct {
	ID   int
	Name string
}

type MileStone struct {
	ID             int64
	ProjectID      int64
	Name           string
	Description    string
	StartDate      time.Time
	ReleaseDueDate time.Time
	Archived       bool
}

type Version struct {
	ID             int64
	ProjectID      int64
	Name           string
	Description    string
	StartDate      time.Time
	ReleaseDueDate time.Time
	Archived       bool
}

type Attachment struct {
	ID          int64
	Name        string
	Size        int64
	createdUser User
	Created     time.Time
}

type Star struct {
	ID        int64
	Comment   string
	URL       string
	Title     string
	Presenter User
	created   time.Time
}

type Category struct {
	ID   int64
	Name string
}

func (b *BackLog) GetIssueFromKey(key string) (*Issue, error) {
	endpoint := "issues/" + key
	backlogURL := fmt.Sprintf(b.baseURL, endpoint)
	var m Issue
	err := getMappedStruct(backlogURL, &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (b *BackLog) GetIssueFromID(id int) (*Issue, error) {
	endpoint := "issues/" + strconv.Itoa(id)
	backlogURL := fmt.Sprintf(b.baseURL, endpoint)
	var m Issue
	err := getMappedStruct(backlogURL, &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// GetIssueURL returns argement issue URL.
func (i *Issue) GetIssueURL(b *BackLog) string {
	URL := fmt.Sprintf("https://%v.backlog.jp/view/%v", b.space, i.IssueKey)
	return URL
}
