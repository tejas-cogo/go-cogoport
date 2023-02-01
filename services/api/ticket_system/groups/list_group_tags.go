package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func ListGroupTag(Tag string) ([]string, error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var t []string

	tx = tx.Table("(?) as u", db.Model(&models.Group{}).Select("unnest(tags) as tag")).Distinct("u.tag")

	if Tag != "" {
		Tag = "%" + Tag + "%"
		db = db.Where("u.tag iLIKE ?", Tag)
	}

	tx.Pluck("tag", &t)

	tx.Commit()
	return t, err
}
