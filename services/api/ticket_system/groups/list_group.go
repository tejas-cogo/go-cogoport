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
		filters.Name = "%" + filters.Name + "%"
		db = db.Where("name Like ?", filters.Name)
	}

	if tags != "" {
		db = db.Where("? Like ANY(tags)", tags)
	}

	if filters.Status != "" {
		db.Where("status = ?", filters.Status)
	}
	db.Order("created_at desc")
	db = db.Find(&groups)

	return groups, db
}
