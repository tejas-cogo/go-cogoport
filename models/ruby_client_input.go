package models

import "github.com/lib/pq"

type RubyClientInput struct {
	Endpoint string
}

type Body struct {
	RmMappingDataRequired bool    `json:"rm_mappings_data_required"`
	Filters               Filters `json:"filters"`
}

type Filters struct {
	RoleIDs pq.StringArray `json:"role_ids"`
	Status  string         `json:"status"`
	UserID  string         `json:"user_id"`
}
