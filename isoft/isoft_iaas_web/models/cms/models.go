package cms

import "time"

type Configuration struct {
	Id                 int64            `json:"id"`                         // 配置项 id
	ParentId           int64            `json:"parent_id"`                  // 父配置项 id,顶级配置为 0
	ConfigurationName  string           `json:"configuration_name"`         // 配置项名称
	ConfigurationValue string           `json:"configuration_value"`        // 配置项值
	SubConfigurations  []*Configuration `json:"sub_configurations" orm:"-"` // 自配置项列表
	CreatedBy          string           `json:"created_by"`
	CreatedTime        time.Time        `json:"created_time"`
	LastUpdatedBy      string           `json:"last_updated_by"`
	LastUpdatedTime    time.Time        `json:"last_updated_time"`
	Status             int              `json:"status"` // 状态 -1 表示失效
}

// 友情链接
type FrindLink struct {
	Id                 int64            `json:"id"`
	LinkName 		   string			`json:"link_name"`
	LinkAddr		   string			`json:"link_addr"`
	CreatedBy          string           `json:"created_by"`
	CreatedTime        time.Time        `json:"created_time"`
	LastUpdatedBy      string           `json:"last_updated_by"`
	LastUpdatedTime    time.Time        `json:"last_updated_time"`
}
