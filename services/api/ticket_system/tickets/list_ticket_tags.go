package ticket_system

import (
	"errors"
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func ListTicketTag(Tag string) ([]string, error) {
	db := config.GetDB()

	var err error

	var t []string

	db = db.Table("(?) as u", db.Model(&models.Ticket{}).Select("unnest(tags) as tag")).Distinct("u.tag")

	if Tag != "" {
		Tag = "%" + Tag + "%"
		db = db.Where("u.tag = ?", Tag)
	}

	if err := db.Pluck("tag", &t).Error; err != nil {
		db.Rollback()
		return t, errors.New("Error Occurred!")
	}


	return t, err
}
