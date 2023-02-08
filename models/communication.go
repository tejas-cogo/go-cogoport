package models

import "github.com/lib/pq"

type Communication struct {
	Recipient    string         `json:"recipient"`
	Type         string         `json:"type"`
	Service      string         `json:"service"`
	ServiceID    string         `json:"service_id"`
	TemplateName string         `json:"template_name"`
	Sender       string         `json:"sender"`
	CcEmails     pq.StringArray `gorm:"type:json[]"`
}
