package models

type RubyClientOutput struct {
	List       []PartnerUserList
	Page       uint
	Total      uint
	TotalCount uint
	PageLimit  uint
}

type PartnerUserList struct {
	UserID string `json:"user_id"`
	// Max    int64
}
