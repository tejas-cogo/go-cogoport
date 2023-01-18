package ticket_system

import (
	"fmt"

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
		db = db.Where("status = ?", filters.Status)
	}

	db = db.Order("name desc").Find(&groups)

	fmt.Println(db.Statement)

	return groups, db
}
