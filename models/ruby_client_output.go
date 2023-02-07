package models

type RubyClientOutput struct {
	List       []PartnerUserList
	Page       uint `json:"page"`
	Total      uint `json:"total"`
	TotalCount uint `json:"total_count"`
	PageLimit  uint `json:"page_limit"`
}

type PartnerUserList struct {
	UserID string `json:"user_id"`
	// Max    int64
}
