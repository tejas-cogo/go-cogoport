package ticket_system

import (
	"strings"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"gorm.io/gorm"
)

func ListGroup(filters models.Group) ([]models.Group, *gorm.DB) {
	db := config.GetDB()

	var groups []models.Group

	if filters.Name != "" {
		filters.Name = "%" + filters.Name + "%"
		db = db.Where("name Like ?", filters.Name)
	}

	if len(filters.Tags) != 0 {
		db = db.Where("tags && ?", "{"+strings.Join(filters.Tags, ",")+"}")
	}

	if filters.Status != "" {
		db = db.Where("status = ?", filters.Status)
	}

	db = db.Order("name desc").Find(&groups)

	return groups, db
}
