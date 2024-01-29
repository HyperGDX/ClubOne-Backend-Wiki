package types

import (
	"time"
)

// type JSONTime time.Time

// func (t JSONTime) MarshalJSON() ([]byte, error) {
// 	//do your serializing here
// 	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("Mon Jan _2"))
// 	return []byte(stamp), nil
// }

type WikiEntry struct {
	Id         int
	Creator    int
	CreateTime time.Time
	UpdateTime time.Time
	Title      string
	Content    string
	Uuid       string
}

type GetWikiEntryResponse struct {
	Creator    int
	Title      string
	Content    string
	CreateTime time.Time
	UpdateTime time.Time
	Uuid       string
}

type CreateWikiEntryRequest struct {
	Creator int
}

type CreateWikiEntryResponse struct {
	Creator    int
	Title      string
	Content    string
	CreateTime time.Time
	UpdateTime time.Time
	Uuid       string
}

type UpdateWikiEntryRequest struct {
	WikiContent string
}

type UpdateWikiEntryResponse struct {
}

type WikiTitle struct {
	Title string `gorm:"column:title"`
	UUID  string `gorm:"column:uuid"`
}
