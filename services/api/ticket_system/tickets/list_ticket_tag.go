package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func ListTicketTag(tag string) []string {
	db := config.GetDB()

	var tickets []models.Ticket

	var tags []string

	if tag != "" {
		db = db.Where("? Like ANY(tags)", tag)
	}

	db.Find(&tickets).Distinct("tags").Pluck("tags", &tags)

	return tags
}
