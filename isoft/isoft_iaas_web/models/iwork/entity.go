package iwork

import "time"

type Entity struct {
	Id              int64     `json:"id"`
	EntityName      string    `json:"entity_name"`
	EntityFieldStr  string    `json:"entity_field_str"`
	CreatedBy       string    `json:"created_by"`
	CreatedTime     time.Time `json:"created_time" orm:"auto_now_add;type(datetime)"`
	LastUpdatedBy   string    `json:"last_updated_by"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
}

