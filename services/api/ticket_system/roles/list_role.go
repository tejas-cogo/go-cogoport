package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"gorm.io/gorm"
)

func ListRole(filters models.Role) ([]models.Role, *gorm.DB) {
	db := config.GetDB()

	var role []models.Role

	if filters.Name != "" {
		filters.Name = "%" + filters.Name + "%"
		db = db.Where("name Like ?", filters.Name)
	}

	if filters.Status != "" {
		db = db.Where("status = ?", filters.Status)
	}

	db.Order("created_at desc")

	db = db.Find(&role)

	return role, db
}
