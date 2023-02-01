package ticket_system

import (
	"errors"
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func ListTicketTag(Tag string) ([]string, error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var t []string

	tx = tx.Table("(?) as u", tx.Model(&models.Ticket{}).Select("unnest(tags) as tag")).Distinct("u.tag")

	if Tag != "" {
		Tag = "%" + Tag + "%"
		tx = tx.Where("u.tag = ?", Tag)
	}

	if err := tx.Pluck("tag", &t).Error; err != nil {
		tx.Rollback()
		return t, errors.New("Error Occurred!")
	}

	tx.Commit()
	return t, err
}
