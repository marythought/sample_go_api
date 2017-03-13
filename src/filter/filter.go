package filter

import "time"

type Filter struct {
	FilterId     int       `json:"filterId"`
	OrgId        int       `json:"orgId"`
	Name         string    `json:"name"`
	Json         string    `json:"json"`
	Type         int       `json:"type"`
	CreatedTime  time.Time `json:"createdTime"`
	ModifiedTime time.Time `json:"modifiedTime"`
}

type Filters []Filter
