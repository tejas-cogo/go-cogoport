package models

type RestClientOutput struct {
	List       []List
	Page       uint `json:"page"`
	Total      uint `json:"total"`
	TotalCount uint `json:"total_count"`
	PageLimit  uint `json:"page_limit"`
}
type List struct {
	UserID             string `json:"user_id"`
	ReportingManagerID string `json:"reporting_manager_id"`
}
