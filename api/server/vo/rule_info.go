package vo

import "time"

type RuleInfo struct {
	Id         uint      `json:"rule_id"`
	Datasource string    `json:"datasource"`
	Condition  string    `json:"condition"`
	Action     string    `json:"action"`
	Owner      uint		 `json:"owner"`
	ParentId   uint	     `json:"belong_company"`
	Stat       string    `json:"stat"`
	CreateTime time.Time `json:"create_time"`
}
