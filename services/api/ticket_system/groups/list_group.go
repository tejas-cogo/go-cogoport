package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"gorm.io/gorm"
)

func ListGroup(filters models.Group, tags string) ([]models.Group, *gorm.DB) {
	db := config.GetDB()

	var groups []models.Group

	if filters.Name != "" {
		db.Where("name Like", filters.Name)
	}

	if tags != "" {
		db.Where("? Like ANY(tags)", tags)
	}

	if filters.Status != "" {
		db.Where("status = ?", filters.Status)
	} else {
		db.Where("status = ?", "active")
	}

	db = db.Find(&groups)

	return groups, db
}
