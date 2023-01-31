package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"gorm.io/gorm"
	"errors"
)

func ListRole(filters models.Role) ([]models.Role, *gorm.DB, error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var role []models.Role

	if filters.Name != "" {
		filters.Name = "%" + filters.Name + "%"
		tx = tx.Where("name iLike ?", filters.Name)
	}

	if filters.Level > 0 {
		tx = tx.Where("level = ?", filters.Level)
	}

	if filters.Status != "" {
		tx = tx.Where("status = ?", filters.Status)
	}


	tx = tx.Order("created_at desc").Find(&role)
	if err := tx.Error; err != nil {
		tx.Rollback()
		return role, tx, errors.New("Error Occurred!")
	}


	tx.Commit()
	return role, tx, err
}
