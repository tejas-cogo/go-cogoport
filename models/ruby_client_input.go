package models

import "github.com/lib/pq"

type RubyClientInput struct {
	Endpoint string
}

type PartnerUserFilter struct {
	RoleIDs pq.StringArray `gorm:"type:json[]"`
	Status  string
}

type PartnerUserBody struct {
	RmMappingDataRequired bool
	Filters               PartnerUserFilter
	PageLimit             uint
}
