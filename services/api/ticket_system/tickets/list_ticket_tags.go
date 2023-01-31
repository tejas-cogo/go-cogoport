package ticket_system

import (
	"fmt"
	"errors"
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"gorm.io/gorm"
)

func ListTicketTag(Tag string) ([]string, *gorm.DB, error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var t []string

	tx = tx.Table("(?) as u", tx.Model(&models.Ticket{}).Select("unnest(tags) as tag")).Distinct("u.tag")

	fmt.Println(Tag)
	if Tag != "" {
		Tag = "%" + Tag + "%"
		tx = tx.Where("u.tag = ?", Tag)
	}

	if err := tx.Pluck("tag", &t).Error; err != nil {
		tx.Rollback()
		return t, tx, errors.New("Error Occurred!")
	}

	tx.Commit()
	return t, tx, err
}
