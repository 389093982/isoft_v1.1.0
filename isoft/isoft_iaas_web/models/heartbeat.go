package models

import (
	"time"
)

type ConfigFile struct {
	Id              int64     `json:"id"`
	EnvInfo         *EnvInfo  `json:"env_info" orm:"rel(fk)"`
	EnvProperty     string    `json:"env_property"`
	EnvValue        string    `json:"env_value"`
	CreatedBy       string    `json:"created_by"`
	CreatedTime     time.Time `json:"created_time"`
	LastUpdatedBy   string    `json:"last_updated_by"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
}

