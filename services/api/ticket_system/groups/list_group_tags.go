package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"gorm.io/gorm"
	"errors"
)

func ListGroupTag(Tag string) ([]string, *gorm.DB, error) {
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
	if err := tx.Error; err != nil {
		tx.Rollback()
		return t, tx, errors.New("Error Occurred!")
	}

	tx.Commit()
	return t, tx, err
}
