package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

type Tags struct{
	Tag []string 
}

func ListTicketTag(tag string) Tags {
	db := config.GetDB()

	var tickets []models.Ticket
	var t Tags

	var tags []string

	if tag != "" {
		db = db.Where("? Like ANY(tags)", tag)
	}

	db.Find(&tickets).Distinct("tags").Pluck("tags", &tags)

	t.Tag = tags
	 
	return t
}
