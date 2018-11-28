package backlog

import "time"

type Space struct {
	SpaceKey           string
	Name               string
	OwnerID            int64
	Lang               string
	Timezone           string
	ReportSendTime     string
	TextFormattingRule string
	Created            time.Time
	Updated            time.Time
}

type SpaceNotification struct {
	Content string
	Updated time.Time
}

type DiskUsage struct {
	Capacity   int64
	Issue      int64
	Wiki       int64
	File       int64
	Subversion int64
	Git        int64
	Details    []DiskUsageDetail
}

type DiskUsageDetail struct {
	ProjectID  int64
	Issue      int64
	Wiki       int64
	File       int64
	Subversion int64
	Git        int64
}

func (b *BackLog) GetSpace() Space {
	return Space{}
}

func (b *BackLog) GetSpaceImage() interface{} {
	return 1
}
