package models

import "github.com/google/uuid"

type RubyClientOutput struct {
	List       []PartnerUserList
	Page       uint
	Total      uint
	TotalCount uint
	PageLimit  uint
}

type PartnerUserList struct {
	ID        string
	PartnerID uuid.UUID 
}
