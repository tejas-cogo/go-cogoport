package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func UpdateGroup(body models.Group) (string, error, models.Group) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var group models.Group

	if body.ID != 0 {
		tx.Where("id = ?", body.ID).Find(&group)
	}

	if body.Name != "" {
		group.Name = body.Name
	}
	if body.Tags != nil {
		group.Tags = body.Tags
	}
	if body.Status != "" {
		group.Status = body.Status
	}

	if err := tx.Save(&group).Error; err != nil {
		tx.Rollback()
		return "Error Occurred!", err, body
	}

	tx.Commit()
	
	return "Successfully Updated!", err, group
}
