package api

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"gorm.io/gorm"
)

func ListRole(filters models.Role) ([]models.Role, *gorm.DB, error) {
	db := config.GetDB()

	var err error

	var role []models.Role

	if filters.Name != "" {
		filters.Name = "%" + filters.Name + "%"
		db = db.Where("name iLike ?", filters.Name)
	}

	if filters.Level > 0 {
		db = db.Where("level = ?", filters.Level)
	}

	if filters.Status != "" {
		db = db.Where("status = ?", filters.Status)
	}

	db = db.Order("created_at desc").Find(&role)


	return role, db, err
}
