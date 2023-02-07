package models

import "github.com/lib/pq"

type RubyClientInput struct {
	Endpoint string
}

type PartnerUserFilter struct {
	RoleIDs pq.StringArray `json:"role_ids"`
	Status  string         `json:"status"`
}

type PartnerUserBody struct {
	RmMappingDataRequired bool              `json:"rm_mappings_data_required"`
	Filters               PartnerUserFilter `json:"filters"`
	PageLimit             uint              `json:"page_limit"`
}
