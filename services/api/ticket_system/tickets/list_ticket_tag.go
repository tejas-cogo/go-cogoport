package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func ListTicketTag(tag string) []string {
	db := config.GetDB()

	var t []string

	if tag != "" {
		db = db.Where("? Like ANY(tags)", tag)
	}

	db.Table("(?) as u", db.Model(&models.Ticket{}).Select("unnest(tags) as tag")).Distinct("u.tag").Pluck("tag", &t)

	return t
}
