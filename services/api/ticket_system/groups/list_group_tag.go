package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func ListGroupTag(Tag string) []string {
	db := config.GetDB()

	var t []string

	db = db.Table("(?) as u", db.Model(&models.Group{}).Select("unnest(tags) as tag")).Distinct("u.tag")

	if Tag != "" {
		Tag = "%" + Tag + "%"
		db = db.Where("u.tag LIKE ?", Tag)
	}

	db.Pluck("tag", &t)

	return t
}
