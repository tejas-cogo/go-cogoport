package ticket_system

import (
	"fmt"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func ListTicketTag(Tag string) []string {
	db := config.GetDB()

	var t []string

	db = db.Table("(?) as u", db.Model(&models.Ticket{}).Select("unnest(tags) as tag")).Distinct("u.tag")

	fmt.Println(Tag)
	if Tag != "" {
		db = db.Where("u.tag = ?", Tag)
	}

	db.Pluck("tag", &t)

	return t
}
