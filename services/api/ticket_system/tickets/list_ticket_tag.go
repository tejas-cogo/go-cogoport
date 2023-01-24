package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func ListTicketTag(tag string) []string {
	db := config.GetDB()

	var t []string

	db = db.Table("(?) as u", db.Model(&models.Ticket{}).Select("unnest(tags) as tag")).Distinct("u.tag")

	if tag != "" {
		db = db.Where("u.tag = ?", tag)
	}

	db.Pluck("tag", &t)

	return t
}
